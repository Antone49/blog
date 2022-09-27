<template>
<div>
    <SidebarVue />
    <DashboardVue :posts="posts"/>
</div>
</template>

<script>
import DashboardVue from "../../components/admin/Dashboard.vue";
import SidebarVue from "../../components/admin/Sidebar.vue";

import {
    getAllPosts,
} from '/functions/post.js'

export default {
    name: "DashboardPage",
    head() {
        return {
            title: 'Dashboard',
        }
    },
    components: {
        SidebarVue,
        DashboardVue,
    },
    data: () => ({
        posts: null
    }),
    mounted: function () {
        getAllPosts().then(result => {
            if (result != null) {
                this.posts = result
            }
        });
    },
};
</script>

<style>
</style>
