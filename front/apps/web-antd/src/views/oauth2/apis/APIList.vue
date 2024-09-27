<script lang="ts" setup>
import type { APIInfo } from '#/types/apis';

import { computed, h, onMounted, ref } from 'vue';
import { usePagination } from 'vue-request';

import {
  DeleteOutlined,
  EditOutlined,
  MoreOutlined,
  RollbackOutlined,
  SearchOutlined,
} from '@ant-design/icons-vue';
import { Divider, Modal } from 'ant-design-vue';

import { DeleteAPI, GetAPIPage, SaveAPI } from '#/store/apis';

import APIForm from './APIForm.vue';

const columns = [
  {
    title: 'Id',
    dataIndex: 'Id',
    key: 'Id',
    sorter: true,
    defaultSortOrder: 'ascend',
  },
  {
    title: 'Enabled',
    dataIndex: 'Enabled',
    key: 'Enabled',
    sorter: true,
  },
  {
    title: 'Name',
    dataIndex: 'Name',
    key: 'Name',
    name: 'Name',
    sorter: true,
    filtered: true,
  },
  {
    title: 'DisplayName',
    dataIndex: 'DisplayName',
    key: 'DisplayName',
    sorter: true,
    filtered: true,
  },
  {
    title: 'Description',
    dataIndex: 'Description',
    key: 'Description',
    filtered: true,
  },
  {
    title: 'Actions',
    key: 'action',
  },
];

const sortField = ref('Id');
const sortOrder = ref('asc');
const searchModel = ref({
  Name: '',
  DisplayName: '',
});
const filters: Array<string> = computed(() => {
  const tmp: Array<string> = [];
  if (searchModel.value.Name !== '') {
    tmp.push(`Name like ?`);
  }
  if (searchModel.value.DisplayName !== '') {
    tmp.push(`Display_Name like ?`);
  }
  return tmp;
});
const args: Array<string> = computed(() => {
  const tmp = [];
  if (searchModel.value.Name !== '') {
    tmp.push(`%${searchModel.value.Name}%`);
  }
  if (searchModel.value.DisplayName !== '') {
    tmp.push(`%${searchModel.value.DisplayName}%`);
  }
  return tmp;
});
const {
  data: dataSource,
  run,
  loading,
  total,
  current,
  pageSize,
} = usePagination(GetAPIPage, {
  defaultParams: {
    page: 1,
    pageSize: 10,
    sortField: 'Id',
    sortOrder: 'asc',
    filters,
    args,
  },
  pagination: {
    currentKey: 'page',
    pageSizeKey: 'pageSize',
  },
});
const fieldMap = {
  Id: 'Id',
  DisplayName: 'display_name',
};

const pagination = computed(() => ({
  total: total.value,
  current: current.value,
  pageSize: pageSize.value,
}));
const searchForm = ref();

const reloadTable = () => {
  setTimeout(() => {
    run({
      page: current.value,
      pageSize: pageSize.value,
      sortField: sortField.value,
      sortOrder: sortOrder.value,
      filters,
      args,
    });
  }, 500);
};
const handleSearch = () => {
  console.log('searchModel', searchModel.value);
  reloadTable();
};
const handleReset = () => {
  searchModel.value = {
    Name: '',
    DisplayName: '',
  };
  searchForm.value.resetFields();
  reloadTable();
};
const open = ref<boolean>(false);
const row = ref<APIInfo>();
const modalForm = ref();
const handleEdit = (record) => {
  row.value = record;

  open.value = true;
};
const handleDelete = (Id) => {
  Modal.confirm({
    title: `Deleting API ${Id}`,
    //  icon: ExclamationCircleOutlined,
    content: `Are you sure delete this API ${Id}?`,
    onOk() {
      DeleteAPI(Id);

      reloadTable();
    },
    onCancel() {
      console.log('Cancel');
    },
  });
  console.log('record', Id);
};

const handleOk = async () => {
  try {
    const v = await modalForm.value.validate();
    if (!v) {
      return;
    }
    console.log('form validate', v);
  } catch (error) {
    console.log('error submit:', error);
    return;
  }
  SaveAPI(row.value);
  open.value = false;
  reloadTable();
};
const handleCancel = () => {
  modalForm.value.resetForm();
  open.value = false;
};
const handleNew = () => {
  row.value = {
    Id: 0,
    Enabled: true,
    Name: '',
    DisplayName: '',
    Description: '',
  };
  open.value = true;
};
const onTableChange = (pagination: any, filters: any, sorters: any) => {
  current.value = pagination.current;
  pageSize.value = pagination.pageSize;
  if (sorters.field) {
    const fieldName = fieldMap[sorters.field];
    sortField.value = fieldName ?? sorters.field;

    sortOrder.value = sorters.order;
  } else {
    sortField.value = 'Id';
    sortOrder.value = 'asc';
  }
  console.log('filter', filters);

  run({
    page: current.value,
    pageSize: pageSize.value,
    sortField: sortField.value,
    sortOrder: sortOrder.value,
    filter: filters,
  });
};

onMounted(() => {
  reloadTable();
});
</script>
<template>
  <div class="p-5">
    <a-page-header
      style="border: 1px solid rgb(235 237 240)"
      sub-title="API list page"
      title="APIs"
    />
    <a-form ref="searchForm" :model="searchModel" layout="inline">
      <a-form-item label="Name">
        <a-input v-model:value="searchModel.Name" placeholder="Name" />
      </a-form-item>
      <a-form-item label="DisplayName">
        <a-input
          v-model:value="searchModel.DisplayName"
          placeholder="DisplayName"
        />
      </a-form-item>
      <a-form-item>
        <a-button type="primary" @click="handleSearch">
          <template #icon><SearchOutlined /></template>
          Search
        </a-button>
        <a-button
          :icon="h(RollbackOutlined)"
          type="primary"
          @click="handleReset"
        >
          Reset
        </a-button>
        <Divider type="vertical" />
        <a-button :icon="h(MoreOutlined)" type="primary" @click="handleNew">
          New
        </a-button>
      </a-form-item>
    </a-form>
    <a-divider type="horizontal" />
    <a-table
      :columns="columns"
      :data-source="dataSource ? dataSource.list : null"
      :loading="loading"
      :pagination="pagination"
      :run="run"
      @change="onTableChange"
    >
      <template #bodyCell="{ column, record }">
        <template v-if="column.key === 'action'">
          <a-space>
            <a-button
              :icon="h(EditOutlined)"
              type="primary"
              @click="handleEdit(record)"
            >
              Edit
            </a-button>
            <Divider type="vertical" />
            <a-button
              :icon="h(DeleteOutlined)"
              danger
              @click="handleDelete(record.Id)"
            >
              Delete
            </a-button>
          </a-space>
        </template>
        <template v-else-if="column.key === 'Enabled'">
          <a-checkbox v-model:checked="record.Enabled" />
        </template>
      </template>
    </a-table>

    <a-modal
      v-model:open="open"
      title="Edit Client Info"
      @cancel="handleCancel"
      @ok="handleOk"
    >
      <APIForm ref="modalForm" :form-model="row" />
    </a-modal>
  </div>
</template>
