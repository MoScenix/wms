<template>
  <a-layout-header class="header">
    <a-row :wrap="false">
      <a-col flex="220px">
        <RouterLink to="/">
          <div class="header-left">
            <img class="logo" src="@/assets/logo.png" alt="Logo" />
            <h1 class="site-title">WMS 仓储管理</h1>
          </div>
        </RouterLink>
      </a-col>
      <a-col flex="auto">
        <a-menu v-model:selectedKeys="selectedKeys" mode="horizontal" :items="menuItems" @click="handleMenuClick" />
      </a-col>
      <a-col>
        <div class="user-login-status">
          <div v-if="loginUserStore.loginUser.id">
            <a-dropdown>
              <a-space>
                <a-avatar :src="loginUserStore.loginUser.userAvatar" />
                {{ loginUserStore.loginUser.userName ?? '无名' }}
              </a-space>
              <template #overlay>
                <a-menu>
                  <a-menu-item @click="router.push('/user/center')">
                    <UserOutlined />
                    个人中心
                  </a-menu-item>
                  <a-menu-divider />
                  <a-menu-item @click="doLogout" danger>
                    <LogoutOutlined />
                    退出登录
                  </a-menu-item>
                </a-menu>
              </template>
            </a-dropdown>
          </div>
          <div v-else>
            <a-button type="primary" href="/user/login">登录</a-button>
          </div>
        </div>
      </a-col>
    </a-row>
  </a-layout-header>
</template>

<script setup lang="ts">
import { computed, h, ref } from 'vue'
import { useRouter } from 'vue-router'
import { type MenuProps, message } from 'ant-design-vue'
import { useLoginUserStore } from '@/stores/loginUser.ts'
import { userLogout } from '@/api/userController.ts'
import { HomeOutlined, LogoutOutlined, UserOutlined } from '@ant-design/icons-vue'

const loginUserStore = useLoginUserStore()
const router = useRouter()
const selectedKeys = ref<string[]>(['/'])

router.afterEach((to) => {
  selectedKeys.value = [to.path]
})

const originItems = [
  {
    key: '/',
    icon: () => h(HomeOutlined),
    label: '主页',
    title: '主页',
  },
  {
    key: '/wms/inventory',
    label: '库存',
    title: '库存查询',
  },
  {
    key: '/wms/inbound',
    label: '入库',
    title: '入库单',
  },
  {
    key: '/wms/outbound',
    label: '出库',
    title: '出库单',
  },
  {
    key: '/wms/txn',
    label: '流水',
    title: '库存流水',
  },
  {
    key: '/wms/warehouse',
    label: '仓库',
    title: '仓库管理',
  },
  {
    key: '/wms/location',
    label: '库位',
    title: '库位管理',
  },
  {
    key: '/wms/sku',
    label: 'SKU',
    title: 'SKU管理',
  },
  {
    key: '/admin/userManage',
    label: '用户管理',
    title: '用户管理',
  },
]

const filterMenus = (menus = [] as MenuProps['items']) => {
  return menus?.filter((menu) => {
    const menuKey = menu?.key as string
    const loginUser = loginUserStore.loginUser
    if (!menuKey) {
      return false
    }
    if (menuKey.startsWith('/admin')) {
      return !!loginUser && loginUser.userRole === 'admin'
    }
    if (menuKey.startsWith('/wms/warehouse') || menuKey.startsWith('/wms/location') || menuKey.startsWith('/wms/sku')) {
      return !!loginUser && loginUser.userRole === 'admin'
    }
    return true
  })
}

const menuItems = computed<MenuProps['items']>(() => filterMenus(originItems))

const handleMenuClick: MenuProps['onClick'] = (e) => {
  const key = e.key as string
  selectedKeys.value = [key]
  if (key.startsWith('/')) {
    router.push(key)
  }
}

const doLogout = async () => {
  const res = await userLogout()
  if (res.data.code === 0) {
    loginUserStore.setLoginUser({
      userName: '未登录',
    })
    message.success('退出登录成功')
    await router.push('/user/login')
  } else {
    message.error('退出登录失败，' + res.data.message)
  }
}
</script>

<style scoped>
.header {
  background: #fff;
  padding: 0 24px;
  border-bottom: 1px solid #e5e9ef;
}

.header-left {
  display: flex;
  align-items: center;
  gap: 12px;
}

.logo {
  height: 42px;
  width: 42px;
}

.site-title {
  margin: 0;
  font-size: 18px;
  color: #1f2a44;
}

.ant-menu-horizontal {
  border-bottom: none !important;
}
</style>
