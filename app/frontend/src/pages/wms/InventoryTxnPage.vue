<template>
  <div id="inventoryTxnPage">
    <a-form layout="inline" :model="searchParams" @finish="doSearch">
      <a-form-item label="业务类型">
        <a-select v-model:value="searchParams.bizType" style="width: 140px" :options="bizTypeOptions" />
      </a-form-item>
      <a-form-item label="业务单号"><a-input v-model:value="searchParams.bizNo" /></a-form-item>
      <a-form-item label="仓库ID"><a-input-number v-model:value="searchParams.warehouseId" /></a-form-item>
      <a-form-item label="SKU ID"><a-input-number v-model:value="searchParams.skuId" /></a-form-item>
      <a-form-item><a-button type="primary" html-type="submit">搜索</a-button></a-form-item>
    </a-form>
    <a-divider />
    <a-table :columns="columns" :data-source="data" :pagination="pagination" @change="doTableChange">
      <template #bodyCell="{ column, record }">
        <template v-if="column.dataIndex === 'bizType'">
          <a-tag :color="record.bizType === 0 ? 'green' : 'orange'">{{ record.bizType === 0 ? '入库' : '出库' }}</a-tag>
        </template>
      </template>
    </a-table>
  </div>
</template>

<script lang="ts" setup>
import { computed, onMounted, reactive, ref } from 'vue'
import { listInventoryTxnByPage } from '@/api/wmsController'

const columns = [
  { title: 'ID', dataIndex: 'id' },
  { title: '业务类型', dataIndex: 'bizType' },
  { title: '业务单号', dataIndex: 'bizNo' },
  { title: '仓库ID', dataIndex: 'warehouseId' },
  { title: '库位ID', dataIndex: 'locationId' },
  { title: 'SKU ID', dataIndex: 'skuId' },
  { title: '批次', dataIndex: 'lotNo' },
  { title: '数量变化', dataIndex: 'qtyChange' },
  { title: '变更前', dataIndex: 'beforeQty' },
  { title: '变更后', dataIndex: 'afterQty' },
  { title: '操作人ID', dataIndex: 'operatorUserId' },
  { title: '时间', dataIndex: 'createTime' },
]
const bizTypeOptions = [
  { label: '入库', value: 0 },
  { label: '出库', value: 1 },
]
const data = ref<any[]>([])
const total = ref(0)
const searchParams = reactive<any>({ pageNum: 1, pageSize: 10, bizType: undefined, bizNo: '', warehouseId: undefined, skuId: undefined })

const fetchData = async () => {
  const res = await listInventoryTxnByPage({ ...searchParams })
  if (res.data.code === 0 && res.data.data) {
    data.value = res.data.data.records || []
    total.value = res.data.data.totalRow || 0
  }
}
const pagination = computed(() => ({ current: searchParams.pageNum, pageSize: searchParams.pageSize, total: total.value, showSizeChanger: true }))
const doTableChange = (page: any) => { searchParams.pageNum = page.current; searchParams.pageSize = page.pageSize; fetchData() }
const doSearch = () => { searchParams.pageNum = 1; fetchData() }
onMounted(fetchData)
</script>

<style scoped>
#inventoryTxnPage { padding: 24px; background: white; margin-top: 16px; }
</style>
