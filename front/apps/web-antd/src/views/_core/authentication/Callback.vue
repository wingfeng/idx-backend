<script type="ts">
import { defineComponent } from 'vue';

import { useAccessStore,useUserStore } from '@vben/stores';

import { UserManager } from 'oidc-client-ts';

import { oidcSetting } from '#/config/oidcSetting';

export default defineComponent({
  name: 'Callback',
  async mounted()  {
    const userManager = new UserManager(oidcSetting)
 //   const page=this
    userManager.signinCallback().then((user) => {
      console.log('user', user)
      console.log('response accessToken',user.access_token)
      useAccessStore().setAccessToken(user.access_token)
      const u = {
        userId: user.profile.sub,
        userName: user.profile.preferred_username,
        email: user.profile.email,
        roles: user.profile.roles,
 //       avatar: `${user.profile.sub}.png`,
      }

      useUserStore().setUserInfo(u)
      console.log('stored userInfo',useUserStore().userInfo)
      console.log('accessToken',useAccessStore().accessToken)
      this.$router.push('/')

    })
  },


})
</script>
<template>
  <div>
    <h1>OIDC Callback Login Successed</h1>
    <RouterLink to="/">Go to Home</RouterLink>
  </div>
</template>
