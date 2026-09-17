<!--此文件在GPL-3.0协议下开源-->
<!--修改者：jackzhang 2026-09-01-->
<!--
 * @Author: liwenjie lwj@holardata.com
 * @Date: 2025-06-16 18:19:16
 * @LastEditors: liwenjie lwj@holardata.com
 * @LastEditTime: 2025-06-16 19:33:28
 * @FilePath: /1Panel/frontend/src/layout/components/AppFooter.vue
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
-->
<template>
    <div class="footer" :style="{ height: mobile ? '108px' : '48px' }">
        <div class="flex w-full flex-col gap-4 md:justify-between md:flex-row">
            <div class="flex flex-wrap gap-4">
                <a href="https://fit2cloud.com/" target="_blank">
                    Copyright © 2014-{{ year }} {{ $t('commons.fit2cloud') }}
                </a>
            </div>
            <div class="flex flex-row gap-2 md:flex-col lg:flex-row">
                <SystemUpgrade :footer="true" />
            </div>
        </div>
    </div>
</template>

<script setup lang="ts">
import { computed } from 'vue';
import SystemUpgrade from '@/components/system-upgrade/index.vue';
import { GlobalStore } from '@/store';

const globalStore = GlobalStore();
const mobile = computed(() => {
    return globalStore.isMobile();
});
const year = new Date().getFullYear();
</script>

<style scoped lang="scss">
.footer {
    display: flex;
    align-items: center;
    justify-content: space-between;
    height: 48px;
    background: var(--panel-footer-bg);
    border-top: 1px solid var(--panel-footer-border);
    box-sizing: border-box;
    padding: 10px 20px;
    a {
        font-size: 12px;
        color: #858585;
        text-decoration: none;
        letter-spacing: 0.5px;
    }
    span {
        font-size: 12px;
        color: #858585;
        text-decoration: none;
        letter-spacing: 0.5px;
    }
}
</style>
