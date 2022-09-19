<template>
<v-container>
    <div class="nav">
        <nuxt-link to="/" exact class="nav__logo">Home</nuxt-link>
        <div class="navbar" :class="{ active_nav: nav }">
            <nuxt-link to="/blog" class="nav_link" exact>Blog</nuxt-link>
            <nuxt-link to="/map" class="nav_link" exact>Map</nuxt-link>
            <nuxt-link to="/about" class="nav_link" exact>About</nuxt-link>
            <nuxt-link to="/contact" class="nav_link" exact>Contact</nuxt-link>
        </div>
    </div>
</v-container>
</template>

<script>
export default {
    name: "NavbarContent",
    data: () => ({
        nav: true,
    }),
};
</script>

<style >
</style>
