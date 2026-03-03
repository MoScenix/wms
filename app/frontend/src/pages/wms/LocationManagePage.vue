<template>
  <div id="locationManagePage">
    <a-form layout="inline" :model="searchParams" @finish="doSearch">
      <a-form-item label="仓库ID"><a-input-number v-model:value="searchParams.warehouseId" style="width: 140px" /></a-form-item>
      <a-form-item label="库位编码"><a-input v-model:value="searchParams.locationCode" /></a-form-item>
      <a-form-item label="库位名称"><a-input v-model:value="searchParams.locationName" /></a-form-item>
      <a-form-item>
        <a-space>
          <a-button type="primary" html-type="submit">搜索</a-button>
          <a-button type="primary" @click="openAdd">新增库位</a-button>
        </a-space>
      </a-form-item>
    </a-form>
    <a-divider />
    <a-table :columns="columns" :data-source="data" :pagination="pagination" @change="doTableChange">
      <template #bodyCell="{ column, record }">
        <template v-if="column.dataIndex === 'createTime'">{{ dayjs(record.createTime).format('YYYY-MM-DD HH:mm:ss') }}</template>
        <template v-else-if="column.key === 'action'"><a-button type="link" @click="openEdit(record)">编辑</a-button></template>
      </template>
    </a-table>

    <a-modal v-model:open="modalOpen" :title="isEdit ? '编辑库位' : '新增库位'" @ok="submitForm">
      <a-form :model="formState" :label-col="{ span: 6 }">
        <a-form-item label="仓库ID" required><a-input-number v-model:value="formState.warehouseId" style="width: 100%" :disabled="isEdit" /></a-form-item>
        <a-form-item label="库位编码" required><a-input v-model:value="formState.locationCode" :disabled="isEdit" /></a-form-item>
        <a-form-item label="库位名称" required><a-input v-model:value="formState.locationName" /></a-form-item>
        <a-form-item label="库位类型"><a-input v-model:value="formState.locationType" /></a-form-item>
      </a-form>
    </a-modal>
  </div>
</template>

<script lang="ts" setup>
import { computed, onMounted, reactive, ref } from 'vue'
import { message } from 'ant-design-vue'
import dayjs from 'dayjs'
import { addLocation, listLocationByPage, updateLocation } from '@/api/wmsController'

const columns = [
  { title: 'ID', dataIndex: 'id' },
  { title: '仓库ID', dataIndex: 'warehouseId' },
  { title: '库位编码', dataIndex: 'locationCode' },
  { title: '库位名称', dataIndex: 'locationName' },
  { title: '库位类型', dataIndex: 'locationType' },
  { title: '创建时间', dataIndex: 'createTime' },
  { title: '操作', key: 'action' },
]
const data = ref<any[]>([])
const total = ref(0)
const searchParams = reactive({ pageNum: 1, pageSize: 10, warehouseId: undefined as number | undefined, locationCode: '', locationName: '' })
const modalOpen = ref(false)
const isEdit = ref(false)
const formState = reactive<any>({ id: undefined, warehouseId: undefined, locationCode: '', locationName: '', locationType: '' })

const fetchData = async () => {
  const res = await listLocationByPage({ ...searchParams })
  if (res.data.code === 0 && res.data.data) {
    data.value = res.data.data.records || []
    total.value = res.data.data.totalRow || 0
  }
}
const pagination = computed(() => ({ current: searchParams.pageNum, pageSize: searchParams.pageSize, total: total.value, showSizeChanger: true }))
const doTableChange = (page: any) => { searchParams.pageNum = page.current; searchParams.pageSize = page.pageSize; fetchData() }
const doSearch = () => { searchParams.pageNum = 1; fetchData() }
const openAdd = () => { isEdit.value = false; Object.assign(formState, { id: undefined, warehouseId: undefined, locationCode: '', locationName: '', locationType: '' }); modalOpen.value = true }
const openEdit = (record: any) => { isEdit.value = true; Object.assign(formState, { ...record }); modalOpen.value = true }
const submitForm = async () => {
  if (!formState.warehouseId || !formState.locationCode || !formState.locationName) { message.warning('请补全必填项'); return }
  const res = isEdit.value
    ? await updateLocation({ id: formState.id, locationName: formState.locationName, locationType: formState.locationType })
    : await addLocation({ warehouseId: formState.warehouseId, locationCode: formState.locationCode, locationName: formState.locationName, locationType: formState.locationType })
  if (res.data.code === 0) { message.success('操作成功'); modalOpen.value = false; fetchData() } else { message.error(res.data.message || '操作失败') }
}
onMounted(fetchData)
</script>

<style scoped>
#locationManagePage { padding: 24px; background: white; margin-top: 16px; }
</style>
