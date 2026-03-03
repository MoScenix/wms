import { createRouter, createWebHistory } from 'vue-router'
import HomePage from '@/pages/HomePage.vue'
import UserLoginPage from '@/pages/user/UserLoginPage.vue'
import UserRegisterPage from '@/pages/user/UserRegisterPage.vue'
import UserManagePage from '@/pages/admin/UserManagePage.vue'
import UserCenterPage from '@/pages/user/UserCenterPage.vue'
import WarehouseManagePage from '@/pages/wms/WarehouseManagePage.vue'
import LocationManagePage from '@/pages/wms/LocationManagePage.vue'
import SkuManagePage from '@/pages/wms/SkuManagePage.vue'
import InventoryPage from '@/pages/wms/InventoryPage.vue'
import InboundOrderPage from '@/pages/wms/InboundOrderPage.vue'
import OutboundOrderPage from '@/pages/wms/OutboundOrderPage.vue'
import InventoryTxnPage from '@/pages/wms/InventoryTxnPage.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: '主页',
      component: HomePage,
    },
    {
      path: '/user/login',
      name: '用户登录',
      component: UserLoginPage,
    },
    {
      path: '/user/register',
      name: '用户注册',
      component: UserRegisterPage,
    },
    {
      path: '/admin/userManage',
      name: '用户管理',
      component: UserManagePage,
    },
    {
      path: '/user/center',
      name: '个人中心',
      component: UserCenterPage,
    },
    {
      path: '/wms/warehouse',
      name: '仓库管理',
      component: WarehouseManagePage,
    },
    {
      path: '/wms/location',
      name: '库位管理',
      component: LocationManagePage,
    },
    {
      path: '/wms/sku',
      name: 'SKU管理',
      component: SkuManagePage,
    },
    {
      path: '/wms/inventory',
      name: '库存查询',
      component: InventoryPage,
    },
    {
      path: '/wms/inbound',
      name: '入库单',
      component: InboundOrderPage,
    },
    {
      path: '/wms/outbound',
      name: '出库单',
      component: OutboundOrderPage,
    },
    {
      path: '/wms/txn',
      name: '库存流水',
      component: InventoryTxnPage,
    },
  ],
})

export default router
