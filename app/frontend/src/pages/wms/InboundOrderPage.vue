<template>
  <div id="inboundOrderPage">
    <a-form layout="inline" :model="searchParams" @finish="doSearch">
      <a-form-item label="单号"><a-input v-model:value="searchParams.inboundNo" /></a-form-item>
      <a-form-item label="仓库ID"><a-input-number v-model:value="searchParams.warehouseId" /></a-form-item>
      <a-form-item label="状态">
        <a-select v-model:value="searchParams.status" style="width: 140px" :options="statusOptions" />
      </a-form-item>
      <a-form-item>
        <a-space>
          <a-button type="primary" html-type="submit">搜索</a-button>
          <a-button type="primary" @click="openCreate">创建入库单</a-button>
        </a-space>
      </a-form-item>
    </a-form>
    <a-divider />
    <a-table :columns="columns" :data-source="data" :pagination="pagination" @change="doTableChange">
      <template #bodyCell="{ column, record }">
        <template v-if="column.dataIndex === 'status'">
          <a-tag :color="record.status === 1 ? 'green' : 'blue'">{{ record.status === 1 ? '已收货' : '草稿' }}</a-tag>
        </template>
        <template v-else-if="column.key === 'action'">
          <a-space>
            <a-button type="link" @click="viewItems(record)">明细</a-button>
            <a-button type="link" :disabled="record.status === 1" @click="doReceive(record.id)">收货</a-button>
          </a-space>
        </template>
      </template>
    </a-table>

    <a-modal v-model:open="createOpen" title="创建入库单" width="760px" @ok="doCreate">
      <a-form :model="createForm" :label-col="{ span: 4 }">
        <a-form-item label="仓库ID" required><a-input-number v-model:value="createForm.warehouseId" style="width: 100%" /></a-form-item>
        <a-form-item label="供应商"><a-input v-model:value="createForm.supplierName" /></a-form-item>
        <a-form-item label="备注"><a-input v-model:value="createForm.remark" /></a-form-item>
        <a-divider>明细</a-divider>
        <div v-for="(item, idx) in createForm.items" :key="idx" style="display: grid; grid-template-columns: repeat(6,1fr) 90px; gap: 8px; margin-bottom: 8px;">
          <a-input-number v-model:value="item.skuId" placeholder="SKU" />
          <a-input-number v-model:value="item.locationId" placeholder="库位" />
          <a-input v-model:value="item.lotNo" placeholder="批次" />
          <a-input v-model:value="item.expireDate" placeholder="效期" />
          <a-input-number v-model:value="item.qty" placeholder="数量" />
          <span></span>
          <a-button danger @click="removeItem(idx)">删除</a-button>
        </div>
        <a-button type="dashed" block @click="addItem">新增明细</a-button>
      </a-form>
    </a-modal>

    <a-modal v-model:open="itemsOpen" title="入库单明细" footer="null">
      <a-table :columns="itemColumns" :data-source="currentItems" :pagination="false" size="small" />
    </a-modal>
  </div>
</template>

<script lang="ts" setup>
import { computed, onMounted, reactive, ref } from 'vue'
import { message } from 'ant-design-vue'
import { createInboundOrder, listInboundOrderByPage, receiveInboundOrder } from '@/api/wmsController'

const columns = [
  { title: 'ID', dataIndex: 'id' },
  { title: '入库单号', dataIndex: 'inboundNo' },
  { title: '仓库ID', dataIndex: 'warehouseId' },
  { title: '供应商', dataIndex: 'supplierName' },
  { title: '状态', dataIndex: 'status' },
  { title: '操作', key: 'action' },
]
const itemColumns = [
  { title: 'SKU', dataIndex: 'skuId' },
  { title: '库位', dataIndex: 'locationId' },
  { title: '批次', dataIndex: 'lotNo' },
  { title: '效期', dataIndex: 'expireDate' },
  { title: '数量', dataIndex: 'qty' },
]
const statusOptions = [
  { label: '草稿', value: 0 },
  { label: '已收货', value: 1 },
]

const data = ref<any[]>([])
const total = ref(0)
const searchParams = reactive<any>({ pageNum: 1, pageSize: 10, inboundNo: '', warehouseId: undefined, status: undefined })
const createOpen = ref(false)
const itemsOpen = ref(false)
const currentItems = ref<any[]>([])
const createForm = reactive<any>({ warehouseId: undefined, supplierName: '', remark: '', items: [{ skuId: undefined, locationId: undefined, lotNo: '', expireDate: '', qty: 1 }] })

const fetchData = async () => {
  const res = await listInboundOrderByPage({ ...searchParams })
  if (res.data.code === 0 && res.data.data) {
    data.value = res.data.data.records || []
    total.value = res.data.data.totalRow || 0
  }
}
const pagination = computed(() => ({ current: searchParams.pageNum, pageSize: searchParams.pageSize, total: total.value, showSizeChanger: true }))
const doTableChange = (page: any) => { searchParams.pageNum = page.current; searchParams.pageSize = page.pageSize; fetchData() }
const doSearch = () => { searchParams.pageNum = 1; fetchData() }
const openCreate = () => {
  Object.assign(createForm, { warehouseId: undefined, supplierName: '', remark: '', items: [{ skuId: undefined, locationId: undefined, lotNo: '', expireDate: '', qty: 1 }] })
  createOpen.value = true
}
const addItem = () => createForm.items.push({ skuId: undefined, locationId: undefined, lotNo: '', expireDate: '', qty: 1 })
const removeItem = (idx: number) => createForm.items.splice(idx, 1)
const doCreate = async () => {
  const res = await createInboundOrder({ ...createForm })
  if (res.data.code === 0) {
    message.success('创建成功')
    createOpen.value = false
    fetchData()
  } else {
    message.error(res.data.message || '创建失败')
  }
}
const doReceive = async (id: number) => {
  const res = await receiveInboundOrder({ id })
  if (res.data.code === 0) {
    message.success('收货成功')
    fetchData()
  } else {
    message.error(res.data.message || '收货失败')
  }
}
const viewItems = (record: any) => {
  currentItems.value = record.items || []
  itemsOpen.value = true
}
onMounted(fetchData)
</script>

<style scoped>
#inboundOrderPage { padding: 24px; background: white; margin-top: 16px; }
</style>
