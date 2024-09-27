<script lang="ts" setup>
import type { Rule } from 'ant-design-vue/es/form';

import type { APIInfo } from '#/types/apis';

import { computed, ref } from 'vue';

const props = defineProps<{
  formModel: APIInfo;
}>();
const model = computed(() => {
  return props.formModel;
});
const form = ref();

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

const labelCol = { style: { width: '150px' } };
const wrapperCol = { span: 14 };
const rules: Record<string, Rule[]> = {
  Name: [
    {
      required: true,
      message: 'Please input API name',
      trigger: 'change',
    },
    { min: 3, max: 25, message: 'Length should be 3 to 25', trigger: 'blur' },
  ],
  DisplayName: [
    {
      required: true,
      message: 'Please input DisplayName',
      trigger: 'change',
    },
    { min: 3, max: 25, message: 'Length should be 3 to 25', trigger: 'blur' },
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

    <a-form-item label="Name" name="Name" required>
      <a-input v-model:value="model.Name" />
    </a-form-item>
    <a-form-item label="DisplayName" name="DisplayName" required>
      <a-input v-model:value="model.DisplayName" />
    </a-form-item>
    <a-form-item label="Description" name="Description" required>
      <a-textarea v-model:value="model.Description" />
    </a-form-item>
  </a-form>
</template>
