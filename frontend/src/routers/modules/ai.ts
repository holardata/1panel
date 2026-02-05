/*
 * @Author: liwenjie lwj@holardata.com
 * @Date: 2025-06-16 18:19:16
 * @LastEditors: liwenjie lwj@holardata.com
 * @LastEditTime: 2026-02-05 14:54:20
 * @FilePath: /1Panel/frontend/src/routers/modules/ai.ts
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
import { Layout } from '@/routers/constant';

const databaseRouter = {
    sort: 4,
    path: '/ai',
    component: Layout,
    redirect: '/ai/model',
    meta: {
        icon: 'p-jiqiren2',
        title: 'menu.ai_tools',
    },
    children: [
        {
            path: '/ai/model',
            name: 'OllamaModel',
            component: () => import('@/views/ai/model/index.vue'),
            meta: {
                title: 'ai_tools.model.model',
                requiresAuth: true,
            },
        },
        {
            path: '/ai/mcp',
            name: 'MCPServer',
            component: () => import('@/views/ai/mcp/server/index.vue'),
            meta: {
                title: 'MCP',
                requiresAuth: true,
            },
        },
        // {
        //     path: '/ai/gpu',
        //     name: 'GPU',
        //     component: () => import('@/views/ai/gpu/index.vue'),
        //     meta: {
        //         title: 'ai_tools.gpu.gpu',
        //         requiresAuth: true,
        //     },
        // },
    ],
};

export default databaseRouter;
