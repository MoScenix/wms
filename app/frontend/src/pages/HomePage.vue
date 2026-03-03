<template>
  <div id="homePage">
    <section class="hero">
      <div class="hero-badge">WMS OPERATIONS HUB</div>
      <h1>仓储作业驾驶舱</h1>
      <p>聚焦主数据、入出库、库存与追溯，按业务链路完成日常作业。</p>
      <div class="hero-actions">
        <a-button type="primary" size="large" @click="router.push('/wms/inbound')">开始入库</a-button>
        <a-button size="large" @click="router.push('/wms/inventory')">查看库存</a-button>
      </div>
    </section>

    <section class="panel">
      <h2>核心功能区</h2>
      <a-row :gutter="16" class="panel-row">
        <a-col :xs="24" :md="12" :lg="8" v-for="item in quickEntrances" :key="item.path">
          <a-card hoverable class="feature-card" @click="router.push(item.path)">
            <template #title>{{ item.title }}</template>
            <p>{{ item.desc }}</p>
            <span class="card-tag">{{ item.group }}</span>
          </a-card>
        </a-col>
      </a-row>
    </section>

    <section class="panel workflow">
      <h2>业务流程</h2>
      <div class="flow-grid">
        <div class="flow-step" v-for="step in workflowSteps" :key="step.title">
          <div class="step-index">{{ step.index }}</div>
          <div class="step-body">
            <h3>{{ step.title }}</h3>
            <p>{{ step.desc }}</p>
          </div>
        </div>
      </div>
    </section>
  </div>
</template>

<script setup lang="ts">
import { computed } from 'vue'
import { useRouter } from 'vue-router'
import { useLoginUserStore } from '@/stores/loginUser'

const router = useRouter()
const loginUserStore = useLoginUserStore()

const quickEntrances = computed(() => {
  const isAdmin = loginUserStore.loginUser?.userRole === 'admin'
  const items = [
    { title: '库存查询', desc: '按仓库、库位、SKU、批次快速定位库存', path: '/wms/inventory', group: '库存' },
    { title: '入库单管理', desc: '创建入库单并执行收货，自动记库存流水', path: '/wms/inbound', group: '单据' },
    { title: '出库单管理', desc: '创建出库单并执行发货，自动扣减可用库存', path: '/wms/outbound', group: '单据' },
    { title: '库存流水', desc: '查询每次库存变化的业务来源与时间', path: '/wms/txn', group: '追溯' },
  ]
  if (isAdmin) {
    items.push(
      { title: '仓库管理', desc: '维护仓库主数据与负责人', path: '/wms/warehouse', group: '主数据' },
      { title: '库位管理', desc: '维护仓内库位编码与类型', path: '/wms/location', group: '主数据' },
      { title: 'SKU 管理', desc: '维护物料编码、名称、单位与分类', path: '/wms/sku', group: '主数据' },
      { title: '用户管理', desc: '维护系统用户角色与权限', path: '/admin/userManage', group: '系统' },
    )
  }
  return items
})

const workflowSteps = [
  { index: '01', title: '准备主数据', desc: '先配置仓库、库位、SKU，确保作业单据有准确基础。' },
  { index: '02', title: '执行入库', desc: '创建并收货入库单，系统自动增加库存并记录流水。' },
  { index: '03', title: '执行出库', desc: '创建并发货出库单，系统校验可用库存后自动扣减。' },
  { index: '04', title: '盘点与追溯', desc: '通过库存查询与流水复盘差异，闭环追踪问题。' },
]
</script>

<style scoped>
#homePage {
  padding: 24px;
  background: linear-gradient(135deg, #edf2f8 0%, #f8fafc 45%, #fff8f0 100%);
  min-height: calc(100vh - 64px - 82px);
}

.hero {
  background: linear-gradient(130deg, #123049 0%, #174d78 55%, #2a6f9f 100%);
  color: #f3f8ff;
  border-radius: 18px;
  padding: 34px 30px;
  box-shadow: 0 16px 30px rgba(17, 48, 73, 0.2);
}

.hero-badge {
  display: inline-block;
  font-size: 12px;
  letter-spacing: 1px;
  background: rgba(255, 255, 255, 0.2);
  padding: 4px 10px;
  border-radius: 999px;
}

.hero h1 {
  margin: 12px 0 8px;
  font-size: 36px;
}

.hero p {
  margin: 0;
  max-width: 640px;
  color: #d8e7f8;
}

.hero-actions {
  margin-top: 22px;
  display: flex;
  gap: 12px;
}

.panel {
  margin-top: 18px;
  padding: 20px;
  border-radius: 14px;
  background: rgba(255, 255, 255, 0.92);
  border: 1px solid #e7edf4;
}

.panel h2 {
  margin-top: 0;
  color: #1f2a44;
}

.feature-card {
  border-radius: 10px;
}

.feature-card p {
  margin-bottom: 12px;
  color: #516078;
}

.card-tag {
  display: inline-block;
  font-size: 12px;
  color: #0d4d76;
  background: #dcefff;
  border-radius: 999px;
  padding: 2px 10px;
}

.workflow {
  margin-bottom: 20px;
}

.flow-grid {
  display: grid;
  grid-template-columns: repeat(2, minmax(0, 1fr));
  gap: 12px;
}

.flow-step {
  display: flex;
  gap: 14px;
  background: #f7fafc;
  border: 1px solid #e4ebf3;
  border-radius: 12px;
  padding: 14px;
}

.step-index {
  font-size: 24px;
  color: #174d78;
  font-weight: 700;
  line-height: 1;
}

.step-body h3 {
  margin: 0 0 6px;
  color: #1f2a44;
}

.step-body p {
  margin: 0;
  color: #516078;
}

@media (max-width: 992px) {
  .flow-grid {
    grid-template-columns: 1fr;
  }

  .hero h1 {
    font-size: 28px;
  }
}
</style>
