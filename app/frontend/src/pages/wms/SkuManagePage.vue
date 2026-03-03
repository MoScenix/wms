<template>
  <div id="skuManagePage">
    <a-form layout="inline" :model="searchParams" @finish="doSearch">
      <a-form-item label="SKU编码"><a-input v-model:value="searchParams.skuCode" /></a-form-item>
      <a-form-item label="SKU名称"><a-input v-model:value="searchParams.skuName" /></a-form-item>
      <a-form-item>
        <a-space>
          <a-button type="primary" html-type="submit">搜索</a-button>
          <a-button type="primary" @click="openAdd">新增SKU</a-button>
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

    <a-modal v-model:open="modalOpen" :title="isEdit ? '编辑SKU' : '新增SKU'" @ok="submitForm">
      <a-form :model="formState" :label-col="{ span: 5 }">
        <a-form-item label="SKU编码" required><a-input v-model:value="formState.skuCode" :disabled="isEdit" /></a-form-item>
        <a-form-item label="SKU名称" required><a-input v-model:value="formState.skuName" /></a-form-item>
        <a-form-item label="单位"><a-input v-model:value="formState.unit" /></a-form-item>
        <a-form-item label="分类"><a-input v-model:value="formState.category" /></a-form-item>
      </a-form>
    </a-modal>
  </div>
</template>

<script lang="ts" setup>
import { computed, onMounted, reactive, ref } from 'vue'
import { message } from 'ant-design-vue'
import dayjs from 'dayjs'
import { addSku, listSkuByPage, updateSku } from '@/api/wmsController'

const columns = [
  { title: 'ID', dataIndex: 'id' },
  { title: 'SKU编码', dataIndex: 'skuCode' },
  { title: 'SKU名称', dataIndex: 'skuName' },
  { title: '单位', dataIndex: 'unit' },
  { title: '分类', dataIndex: 'category' },
  { title: '创建时间', dataIndex: 'createTime' },
  { title: '操作', key: 'action' },
]
const data = ref<any[]>([])
const total = ref(0)
const searchParams = reactive({ pageNum: 1, pageSize: 10, skuCode: '', skuName: '' })
const modalOpen = ref(false)
const isEdit = ref(false)
const formState = reactive<any>({ id: undefined, skuCode: '', skuName: '', unit: '', category: '' })

const fetchData = async () => {
  const res = await listSkuByPage({ ...searchParams })
  if (res.data.code === 0 && res.data.data) {
    data.value = res.data.data.records || []
    total.value = res.data.data.totalRow || 0
  }
}
const pagination = computed(() => ({ current: searchParams.pageNum, pageSize: searchParams.pageSize, total: total.value, showSizeChanger: true }))
const doTableChange = (page: any) => { searchParams.pageNum = page.current; searchParams.pageSize = page.pageSize; fetchData() }
const doSearch = () => { searchParams.pageNum = 1; fetchData() }
const openAdd = () => { isEdit.value = false; Object.assign(formState, { id: undefined, skuCode: '', skuName: '', unit: '', category: '' }); modalOpen.value = true }
const openEdit = (record: any) => { isEdit.value = true; Object.assign(formState, { ...record }); modalOpen.value = true }
const submitForm = async () => {
  if (!formState.skuCode || !formState.skuName) { message.warning('请补全必填项'); return }
  const res = isEdit.value
    ? await updateSku({ id: formState.id, skuName: formState.skuName, unit: formState.unit, category: formState.category })
    : await addSku({ skuCode: formState.skuCode, skuName: formState.skuName, unit: formState.unit, category: formState.category })
  if (res.data.code === 0) { message.success('操作成功'); modalOpen.value = false; fetchData() } else { message.error(res.data.message || '操作失败') }
}
onMounted(fetchData)
</script>

<style scoped>
#skuManagePage { padding: 24px; background: white; margin-top: 16px; }
</style>
