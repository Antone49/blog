<template>
<div>
    <NavbarMobileVue />
    <PostDetailHeroVue :image="image" />
    <PostDetailVue :title="title" :content="content" :tags="tags" />
    <PostMessageVue />
    <FooterVue />
</div>
</template>

<script>
import FooterVue from "/components/front/Footer.vue";
import NavbarMobileVue from "/components/front/NavbarMobile.vue";
import PostDetailVue from "/components/front/PostDetail.vue";
import PostMessageVue from "/components/front/PostMessage.vue";
import PostDetailHeroVue from "/components/front/PostDetailHero.vue";

import {
    getPost,
    getPostTags
} from '/functions/post.js'

export default {
    name: "PostDetail",
    auth: false,
    head() {
        return {
            title: this.title,
        }
    },
    components: {
        NavbarMobileVue,
        FooterVue,
        PostDetailHeroVue,
        PostDetailVue,
        PostMessageVue,
    },
    data: () => ({
        title: null,
        content: null,
        image: "http://localhost:8800/images/loading.png",
        tags: null
    }),
    mounted: function () {
        var idPost = parseInt(this.$route.query.id)

        getPost(idPost).then(result => {
            if (result != null) {
                this.title = result.title
                this.content = result.content
                this.image = "http://localhost:8800" + result.image
            }
        });

        getPostTags(idPost).then(result => {
            if (result != null) {
                this.tags = result
            }
        });
    },
};
</script>

<style>
</style>
