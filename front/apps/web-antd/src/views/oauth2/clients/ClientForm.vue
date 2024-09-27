<script lang="ts" setup>
import type { Rule } from 'ant-design-vue/es/form';

import type { ClientInfo } from '#/types/client';

import { computed, ref } from 'vue';

import { OAuth2GrantTypes } from '#/config/grandtypes';

const props = defineProps<{
  client: ClientInfo;
}>();
const model = computed(() => {
  return props.client;
});
const form = ref();
const scopeOptions = [
  { label: 'openid', value: 'openid' },
  { label: 'profile', value: 'profile' },
  { label: 'email', value: 'email' },
  { label: 'address', value: 'address' },
  { label: 'phone', value: 'phone' },
];
const scopes = computed({
  get() {
    return model.value.Scopes.split(' ');
  },
  set(newVal) {
    model.value.Scopes = newVal.join(' ');
  },
});

const validate = () => {
  return form.value
    .validate()
    .then(() => {
      return true;
    })
    .catch((error: any) => {
      console.log('error submit:', error);
      return false;
    });
};
const resetForm = () => {
  form.value.resetFields();
};
const grantTypesOptions = computed(() => {
  const types = [];
  OAuth2GrantTypes.forEach((key, v) => {
    types.push({ label: key, value: v });
  });

  return types;
});
const grantTypes = computed({
  get() {
    return model.value.GrantTypes.split(' ');
  },
  set(newVal) {
    model.value.GrantTypes = newVal.join(' ');
  },
});
const redirectUris = computed({
  get() {
    const array = model.value.RedirectUris
      ? model.value.RedirectUris.split(' ')
      : [];
    return array.join('\n');
  },
  set(newVal) {
    const array = newVal.split('\n');
    model.value.RedirectUris = array.join(' ');
  },
});
const labelCol = { style: { width: '150px' } };
const wrapperCol = { span: 14 };
const rules: Record<string, Rule[]> = {
  ClientName: [
    {
      required: true,
      message: 'Please input Client name',
      trigger: 'change',
    },
    { min: 3, max: 25, message: 'Length should be 3 to 25', trigger: 'blur' },
  ],
  ClientId: [
    {
      required: true,
      message: 'Please input Client Id',
      trigger: 'change',
    },
    { min: 3, max: 25, message: 'Length should be 3 to 25', trigger: 'blur' },
  ],
  Scopes: [
    {
      required: true,
      message: 'Scopes cannot be empty',
      trigger: 'change',
    },
  ],

  GrantTypes: [
    {
      required: true,
      message: 'GrantTypes cannot be empty',
      trigger: 'change',
    },
  ],
};
defineExpose({ validate, resetForm, model });
</script>
<template>
  <a-form
    ref="form"
    :label-col="labelCol"
    :model="model"
    :rules="rules"
    :wrapper-col="wrapperCol"
  >
    <a-form-item label="Id">
      <a-label>{{ model.Id }}</a-label>
    </a-form-item>
    <a-form-item label="Enabled">
      <a-checkbox v-model:checked="model.Enabled" />
    </a-form-item>

    <a-form-item label="Client name" name="ClientName" required>
      <a-input v-model:value="model.ClientName" />
    </a-form-item>
    <a-form-item label="Client Id" name="ClientId" required>
      <a-input v-model:value="model.ClientId" />
    </a-form-item>
    <a-form-item label=" ">
      <a-checkbox v-model:checked="model.RequireConsent">
        Require consent
      </a-checkbox>
      <a-checkbox v-model:checked="model.RequireSecret">
        Require Secret
      </a-checkbox>
      <a-checkbox v-model:checked="model.RequirePkce">
        Require PKCE
      </a-checkbox>
    </a-form-item>
    <a-form-item label="Client Url">
      <a-input v-model:value="model.ClientURI" />
    </a-form-item>
    <a-form-item label="Client Logo Url">
      <a-image :src="client.LogoURI" height="32px" width="32px" />
      <a-input v-model:value="model.LogoURI" />
    </a-form-item>
    <a-form-item label="Description">
      <a-textarea v-model:value="model.Description" />
    </a-form-item>
    <a-form-item label="Grant types" name="GrantTypes">
      <a-select
        v-model:value="grantTypes"
        :options="grantTypesOptions"
        mode="multiple"
      />
    </a-form-item>
    <a-form-item label="Scopes" name="Scopes">
      <a-select
        v-model:value="scopes"
        :options="scopeOptions"
        mode="multiple"
      />
    </a-form-item>
    <a-form-item label="Redirect Uris">
      <a-textarea v-model:value="redirectUris" />
    </a-form-item>
  </a-form>
</template>
