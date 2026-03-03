<template>
  <div id="inventoryPage">
    <a-form layout="inline" :model="searchParams" @finish="doSearch">
      <a-form-item label="仓库ID"><a-input-number v-model:value="searchParams.warehouseId" /></a-form-item>
      <a-form-item label="库位ID"><a-input-number v-model:value="searchParams.locationId" /></a-form-item>
      <a-form-item label="SKU ID"><a-input-number v-model:value="searchParams.skuId" /></a-form-item>
      <a-form-item label="批次"><a-input v-model:value="searchParams.lotNo" /></a-form-item>
      <a-form-item>
        <a-space>
          <a-button type="primary" html-type="submit">搜索</a-button>
        </a-space>
      </a-form-item>
    </a-form>
    <a-divider />
    <a-table :columns="columns" :data-source="data" :pagination="pagination" @change="doTableChange">
      <template #bodyCell="{ column, record }">
        <template v-if="column.dataIndex === 'createTime'">{{ dayjs(record.createTime).format('YYYY-MM-DD HH:mm:ss') }}</template>
        <template v-else-if="column.dataIndex === 'warning'">
          <a-tag :color="record.availableQuantity <= 0 ? 'red' : 'green'">
            {{ record.availableQuantity <= 0 ? '库存不足' : '正常' }}
          </a-tag>
        </template>
      </template>
    </a-table>
  </div>
</template>

<script lang="ts" setup>
import { computed, onMounted, reactive, ref } from 'vue'
import dayjs from 'dayjs'
import { listInventoryByPage } from '@/api/wmsController'

const columns = [
  { title: 'ID', dataIndex: 'id' },
  { title: '仓库ID', dataIndex: 'warehouseId' },
  { title: '库位ID', dataIndex: 'locationId' },
  { title: 'SKU ID', dataIndex: 'skuId' },
  { title: '批次', dataIndex: 'lotNo' },
  { title: '效期', dataIndex: 'expireDate' },
  { title: '总库存', dataIndex: 'quantity' },
  { title: '锁定库存', dataIndex: 'lockedQuantity' },
  { title: '可用库存', dataIndex: 'availableQuantity' },
  { title: '状态', dataIndex: 'warning' },
  { title: '创建时间', dataIndex: 'createTime' },
]

const data = ref<any[]>([])
const total = ref(0)
const searchParams = reactive<any>({ pageNum: 1, pageSize: 10, warehouseId: undefined, locationId: undefined, skuId: undefined, lotNo: '' })
const fetchData = async () => {
  const res = await listInventoryByPage({ ...searchParams })
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
#inventoryPage { padding: 24px; background: white; margin-top: 16px; }
</style>
