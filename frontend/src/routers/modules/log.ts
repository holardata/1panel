/*
 * @Author: liwenjie lwj@holardata.com
 * @Date: 2025-06-16 18:19:16
 * @LastEditors: liwenjie lwj@holardata.com
 * @LastEditTime: 2025-06-17 19:24:14
 * @FilePath: /1Panel/frontend/src/routers/modules/log.ts
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
import { Layout } from '@/routers/constant';

const logsRouter = {
    sort: 10,
    path: '/logs',
    component: Layout,
    redirect: '/logs/operation',
    meta: {
        title: 'menu.logs',
        icon: 'p-log',
    },
    children: [
        {
            path: '/logs',
            name: 'Log',
            redirect: '/logs/operation',
            component: () => import('@/views/log/index.vue'),
            meta: {},
            children: [
                {
                    path: 'operation',
                    name: 'OperationLog',
                    component: () => import('@/views/log/operation/index.vue'),
                    hidden: true,
                    meta: {
                        activeMenu: '/logs',
                        requiresAuth: false,
                    },
                },
                {
                    path: 'login',
                    name: 'LoginLog',
                    component: () => import('@/views/log/login/index.vue'),
                    hidden: true,
                    meta: {
                        activeMenu: '/logs',
                        requiresAuth: false,
                    },
                },
                {
                    path: 'website',
                    name: 'WebsiteLog',
                    component: () => import('@/views/log/website/index.vue'),
                    hidden: true,
                    meta: {
                        activeMenu: '/logs',
                        requiresAuth: false,
                    },
                },
                {
                    path: 'system',
                    name: 'SystemLog',
                    component: () => import('@/views/log/system/index.vue'),
                    hidden: true,
                    meta: {
                        activeMenu: '/logs',
                        requiresAuth: false,
                    },
                },
                // {
                //     path: 'ssh',
                //     name: 'SSHLog2',
                //     component: () => import('@/views/host/ssh/log/log.vue'),
                //     hidden: true,
                //     meta: {
                //         activeMenu: '/logs',
                //         requiresAuth: false,
                //     },
                // },
            ],
        },
    ],
};

export default logsRouter;
