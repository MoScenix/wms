<template>
  <div id="warehouseManagePage">
    <a-form layout="inline" :model="searchParams" @finish="doSearch">
      <a-form-item label="仓库编码">
        <a-input v-model:value="searchParams.warehouseCode" placeholder="输入仓库编码" />
      </a-form-item>
      <a-form-item label="仓库名称">
        <a-input v-model:value="searchParams.warehouseName" placeholder="输入仓库名称" />
      </a-form-item>
      <a-form-item>
        <a-space>
          <a-button type="primary" html-type="submit">搜索</a-button>
          <a-button type="primary" @click="openAdd">新增仓库</a-button>
        </a-space>
      </a-form-item>
    </a-form>
    <a-divider />
    <a-table :columns="columns" :data-source="data" :pagination="pagination" @change="doTableChange">
      <template #bodyCell="{ column, record }">
        <template v-if="column.dataIndex === 'createTime'">
          {{ dayjs(record.createTime).format('YYYY-MM-DD HH:mm:ss') }}
        </template>
        <template v-else-if="column.key === 'action'">
          <a-button type="link" @click="openEdit(record)">编辑</a-button>
        </template>
      </template>
    </a-table>

    <a-modal v-model:open="modalOpen" :title="isEdit ? '编辑仓库' : '新增仓库'" @ok="submitForm">
      <a-form :model="formState" :label-col="{ span: 5 }">
        <a-form-item label="仓库编码" required>
          <a-input v-model:value="formState.warehouseCode" :disabled="isEdit" />
        </a-form-item>
        <a-form-item label="仓库名称" required>
          <a-input v-model:value="formState.warehouseName" />
        </a-form-item>
        <a-form-item label="地址">
          <a-input v-model:value="formState.address" />
        </a-form-item>
        <a-form-item label="负责人ID">
          <a-input-number v-model:value="formState.managerUserId" style="width: 100%" />
        </a-form-item>
      </a-form>
    </a-modal>
  </div>
</template>

<script lang="ts" setup>
import { computed, onMounted, reactive, ref } from 'vue'
import { message } from 'ant-design-vue'
import dayjs from 'dayjs'
import { addWarehouse, listWarehouseByPage, updateWarehouse } from '@/api/wmsController'

const columns = [
  { title: 'ID', dataIndex: 'id' },
  { title: '仓库编码', dataIndex: 'warehouseCode' },
  { title: '仓库名称', dataIndex: 'warehouseName' },
  { title: '地址', dataIndex: 'address' },
  { title: '负责人ID', dataIndex: 'managerUserId' },
  { title: '创建时间', dataIndex: 'createTime' },
  { title: '操作', key: 'action' },
]

const data = ref<any[]>([])
const total = ref(0)
const searchParams = reactive({ pageNum: 1, pageSize: 10, warehouseCode: '', warehouseName: '' })

const modalOpen = ref(false)
const isEdit = ref(false)
const formState = reactive<any>({ id: undefined, warehouseCode: '', warehouseName: '', address: '', managerUserId: 0 })

const fetchData = async () => {
  const res = await listWarehouseByPage({ ...searchParams })
  if (res.data.code === 0 && res.data.data) {
    data.value = res.data.data.records || []
    total.value = res.data.data.totalRow || 0
  }
}

const pagination = computed(() => ({
  current: searchParams.pageNum,
  pageSize: searchParams.pageSize,
  total: total.value,
  showSizeChanger: true,
  showTotal: (t: number) => `共 ${t} 条`,
}))

const doTableChange = (page: any) => {
  searchParams.pageNum = page.current
  searchParams.pageSize = page.pageSize
  fetchData()
}

const doSearch = () => {
  searchParams.pageNum = 1
  fetchData()
}

const openAdd = () => {
  isEdit.value = false
  Object.assign(formState, { id: undefined, warehouseCode: '', warehouseName: '', address: '', managerUserId: 0 })
  modalOpen.value = true
}

const openEdit = (record: any) => {
  isEdit.value = true
  Object.assign(formState, { ...record })
  modalOpen.value = true
}

const submitForm = async () => {
  if (!formState.warehouseCode || !formState.warehouseName) {
    message.warning('请补全必填项')
    return
  }
  const res = isEdit.value
    ? await updateWarehouse({ id: formState.id, warehouseName: formState.warehouseName, address: formState.address, managerUserId: formState.managerUserId })
    : await addWarehouse({ warehouseCode: formState.warehouseCode, warehouseName: formState.warehouseName, address: formState.address, managerUserId: formState.managerUserId })

  if (res.data.code === 0) {
    message.success('操作成功')
    modalOpen.value = false
    fetchData()
  } else {
    message.error(res.data.message || '操作失败')
  }
}

onMounted(fetchData)
</script>

<style scoped>
#warehouseManagePage {
  padding: 24px;
  background: white;
  margin-top: 16px;
}
</style>
