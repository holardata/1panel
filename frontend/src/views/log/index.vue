<!--
 * @Author: liwenjie lwj@holardata.com
 * @Date: 2025-06-16 18:19:16
 * @LastEditors: liwenjie lwj@holardata.com
 * @LastEditTime: 2025-06-17 19:28:31
 * @FilePath: /1Panel/frontend/src/views/log/index.vue
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
-->
<template>
    <div>
        <RouterButton :buttons="buttons" />
        <LayoutContent>
            <router-view></router-view>
        </LayoutContent>
    </div>
</template>

<script lang="ts" setup>
import i18n from '@/lang';

const buttons = [
    {
        label: i18n.global.t('logs.panelLog'),
        path: '/logs/operation',
    },
    // {
    //     label: 'SSH ' + i18n.global.t('ssh.loginLogs'),
    //     path: '/logs/ssh',
    // },
    // {
    //     label: i18n.global.t('logs.websiteLog'),
    //     path: '/logs/website',
    // },
];
</script>
