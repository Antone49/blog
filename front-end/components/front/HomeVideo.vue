<template>
<div>
    <video-background src="images/boat.mp4" style="max-height: 1200px; height: 100vh;">
        <NavbarContentVue />
        <div class="hero_title">
            <h1>Dhaka</h1>
        </div>
    </video-background>
</div>
</template>

<script>
import NavbarContentVue from "./NavbarContent.vue";

export default {
    name: "HeroCamp",
    components: {
        NavbarContentVue,
        [process.client && 'VideoBackground']: () => import('vue-responsive-video-background-player'),
    },
    data: () => ({
        model: 0,
    }),
};
</script>

<style>
</style>
