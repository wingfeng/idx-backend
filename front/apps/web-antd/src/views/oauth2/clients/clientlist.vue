<script lang="ts" setup>
import type { ClientInfo } from '#/types/client';

import { computed, h, ref } from 'vue';
import { usePagination } from 'vue-request';

import {
  DeleteOutlined,
  EditOutlined,
  ExclamationCircleOutlined,
  MoreOutlined,
  RollbackOutlined,
  SearchOutlined,
} from '@ant-design/icons-vue';
import { Divider, Modal } from 'ant-design-vue';

import { DeleteClient, GetClientsPage, SaveClient } from '#/store/client';

import ClientForm from './ClientForm.vue';

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
    title: 'ClientName',
    dataIndex: 'ClientName',
    key: 'ClientName',
    name: 'client_name',
    sorter: true,
    filtered: true,
  },
  {
    title: 'ClientId',
    dataIndex: 'ClientId',
    key: 'client_id',
    sorter: true,
    filtered: true,
  },
  {
    title: 'Scopes',
    dataIndex: 'Scopes',
    key: 'Scopes',
    sorter: true,
    filtered: true,
  },
  {
    title: 'GrantTypes',
    dataIndex: 'GrantTypes',
    key: 'GrantTypes',
    sorter: true,
    filtered: true,
  },
  {
    title: 'Actions',
    key: 'action',
  },
];
const searchForm = ref();

const searchModel = ref({
  ClientId: '',
  ClientName: '',
});
const sortField = ref('Id');
const sortOrder = ref('asc');
const filters = computed(() => {
  const tmp = [];
  if (searchModel.value.ClientId !== '') {
    tmp.push(`Client_Id like ?`);
  }
  if (searchModel.value.ClientName !== '') {
    tmp.push(`Client_Name like ?`);
  }
  return tmp;
});
const args: Array<string> = computed(() => {
  const tmp = [];
  if (searchModel.value.ClientId !== '') {
    tmp.push(`%${searchModel.value.ClientId}%`);
  }
  if (searchModel.value.ClientName !== '') {
    tmp.push(`%${searchModel.value.ClientName}%`);
  }
  return tmp;
});
const {
  data: dataSource,
  run,
  refresh,
  loading,
  total,
  current,
  pageSize,
} = usePagination(GetClientsPage, {
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
  ClientName: 'Client_Name',
  ClientId: 'Client_Id',
};

const pagination = computed(() => ({
  total: total.value,
  current: current.value,
  pageSize: pageSize.value,
}));

const reloadTable = () => {
  // setTimeout(() => {
  //   run({
  //     page: current.value,
  //     pageSize: pageSize.value,
  //     sortField: sortField.value,
  //     sortOrder: sortOrder.value,
  //     filter,
  //   });
  // }, 500);
  setTimeout(() => {
    refresh();
  }, 50);
};
const handleSearch = () => {
  console.log('searchModel', searchModel.value);
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
const handleReset = () => {
  searchModel.value = {
    ClientId: '',
    ClientName: '',
  };
  searchForm.value.resetFields();
  reloadTable();
};
const open = ref<boolean>(false);
const row = ref<ClientInfo>();
const clientForm = ref();
const handleEdit = (record) => {
  row.value = record;

  open.value = true;
};
const handleDelete = (Id) => {
  Modal.confirm({
    title: `Deleting Client ${Id}`,
    icon: ExclamationCircleOutlined,
    content: `Are you sure delete this client ${Id}?`,
    onOk() {
      DeleteClient(Id);

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
    const v = await clientForm.value.validate();
    if (!v) {
      return;
    }
    console.log('form validate', v);
  } catch (error) {
    console.log('error submit:', error);
    return;
  }
  SaveClient(row.value);
  open.value = false;
  reloadTable();
};
const handleCancel = () => {
  clientForm.value.resetForm();
  open.value = false;
};
const handleNew = () => {
  row.value = {
    Id: 0,
    Enabled: true,
    ClientId: '',
    ClientName: '',
    GrantTypes: 'authorization_code',
    Scopes: 'openid',
    RequireConsent: true,
    RequirePkce: false,
    RequireSecret: true,
    RedirectUris: 'http://localhost:8080/signin-oidc',
  };
  open.value = true;
};
const onTableChange = (pagination: any, filters: any, sorters: any) => {
  current.value = pagination.current;
  pageSize.value = pagination.pageSize;
  if (sorters.field) {
    sortField.value = fieldMap[sorters.field] ?? sorters.field; // 由于数据库字段和Json显示字段不一致，所以做了Mapping
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
// onMounted(() => {
//   reloadTable();
// });
</script>
<template>
  <div class="p-5">
    <a-page-header
      style="border: 1px solid rgb(235 237 240)"
      sub-title="Client list page"
      title="Clients"
    />
    <a-form ref="searchForm" :model="searchModel" layout="inline">
      <a-form-item label="ClientId">
        <a-input v-model:value="searchModel.ClientId" placeholder="ClientId" />
      </a-form-item>
      <a-form-item label="ClientName">
        <a-input
          v-model:value="searchModel.ClientName"
          placeholder="ClientName"
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
      <ClientForm ref="clientForm" :client="row" />
    </a-modal>
  </div>
</template>
