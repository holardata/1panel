<!--
 * @Author: liwenjie lwj@holardata.com
 * @Date: 2025-06-16 18:19:16
 * @LastEditors: liwenjie lwj@holardata.com
 * @LastEditTime: 2025-06-19 16:35:29
 * @FilePath: /1Panel/frontend/src/views/setting/index.vue
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
        label: i18n.global.t('setting.panel'),
        path: '/settings/panel',
    },
    // {
    //     label: i18n.global.t('setting.safe'),
    //     path: '/settings/safe',
    // },
    // {
    //     label: i18n.global.t('setting.backupAccount', 2),
    //     path: '/settings/backupaccount',
    // },
    // {
    //     label: i18n.global.t('setting.snapshot', 2),
    //     path: '/settings/snapshot',
    // },
    // {
    //     label: i18n.global.t('setting.license'),
    //     path: '/settings/license',
    // },
    // {
    //     label: i18n.global.t('setting.about'),
    //     path: '/settings/about',
    // },
];
</script>
