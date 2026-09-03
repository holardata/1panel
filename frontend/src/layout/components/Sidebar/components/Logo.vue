<!--此文件在GPL-3.0协议下开源-->
<!--修改者：jackzhang 2026-09-01-->
<!--
 * @Author: liwenjie lwj@holardata.com
 * @Date: 2025-06-16 18:19:16
 * @LastEditors: liwenjie lwj@holardata.com
 * @LastEditTime: 2025-06-17 11:10:41
 * @FilePath: /1Panel/frontend/src/layout/components/Sidebar/components/Logo.vue
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
-->
<template>
    <div class="logo" style="cursor: pointer" @click="goHome">
        <template v-if="isCollapse">
            <img
                v-if="globalStore.themeConfig.logo"
                :src="`/api/v1/images/logo?t=${Date.now()}`"
                style="cursor: pointer"
                alt="logo"
            />
            <!-- <MenuLogo v-else /> -->
            <img :src="MenuLogo" style="height: 30px; width: auto" v-else />
        </template>
        <template v-else>
            <img
                v-if="globalStore.themeConfig.logoWithText"
                :src="`/api/v1/images/logoWithText?t=${Date.now()}`"
                style="cursor: pointer"
                alt="logo"
            />
            <PrimaryLogo style="width: 90%; height: auto" v-else />
            <!-- <img :src="PrimaryLogo" alt="" /> -->
        </template>
    </div>
</template>

<script setup lang="ts">
import router from '@/routers';
import { GlobalStore } from '@/store';
import PrimaryLogo from '@/assets/images/1panel-logo.svg?component';
import MenuLogo from '@/assets/images/1panel-menu-logo.png';

defineProps<{ isCollapse: boolean }>();

const globalStore = GlobalStore();

const goHome = () => {
    router.push({ name: 'home' });
};
</script>

<style scoped lang="scss">
.logo {
    display: flex;
    align-items: center;
    justify-content: center;
    height: 55px;
    z-index: 1;
    img {
        object-fit: contain;
        width: 95%;
        height: 45px;
    }
}
</style>
