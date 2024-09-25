<script lang="ts">
import { defineComponent } from 'vue';

import { UserManager } from 'oidc-client-ts';

import { oidcSetting } from '#/config/oidcSetting';

export default defineComponent({
  name: 'Login',
  mounted() {
    const userManager = new UserManager(oidcSetting);
    setTimeout(() => {
      userManager.signinRedirect();
    }, 500);
  },
});
</script>

<template>
  <div>
    <h1>IDX Login</h1>
  </div>
</template>
