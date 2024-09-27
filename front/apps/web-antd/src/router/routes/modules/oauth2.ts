import type { RouteRecordRaw } from 'vue-router';

import { BasicLayout } from '#/layouts';

const routes: RouteRecordRaw[] = [
  {
    component: BasicLayout,
    meta: {
      badgeType: 'dot',
      icon: 'ic:baseline-view-in-ar',
      order: 9999,
      title: 'Manage',
    },
    name: 'OAuth2',
    path: '/oauth2-admin',
    children: [
      {
        name: 'Clients',
        path: 'clientlist',
        component: () => import('#/views/oauth2/clients/ClientList.vue'),
        meta: {
          icon: 'lucide:copyright',
          title: 'Clients',
        },
      },
      {
        name: 'APIs',
        path: 'apis',
        component: () => import('#/views/oauth2/apis/APIList.vue'),
        meta: {
          icon: 'lucide:book-open-text',

          title: 'APIs',
        },
      },
      {
        name: 'Organization Unit',
        path: 'ou',
        component: () => import('#/views/oauth2/apis/APIList.vue'),
        meta: {
          icon: 'mdi:github',

          title: 'Organization Unit',
        },
      },
      {
        name: 'Groups',
        path: 'Groups',
        component: () => import('#/views/oauth2/apis/APIList.vue'),
        meta: {
          badgeType: 'dot',

          title: 'Groups',
        },
      },
      {
        name: 'Users',
        path: 'users',
        component: () => import('#/views/oauth2/apis/APIList.vue'),
        meta: {
          badgeType: 'dot',

          title: 'Users',
        },
      },
    ],
  },
];

export default routes;
