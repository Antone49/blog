<template>
<div>
    <NavbarMobileVue />
    <PostDetailHeroVue :title="item.title" />
    <PostDetailVue :category="item.category" :title="item.title" :content="item.content" :image="item.image" :lastestPosts="lastestPosts" :tags="tags" />
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
    getLastestPosts
} from '/functions/post.js'

import {
    getAllTags
} from '/functions/tag.js'

export default {
    name: "PostDetail",
    auth: false,
    head() {
        return {
            title: this.item.title,
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
        item: {
            title: null,
            content: null,
            category: null,
            image: "loading.png",
        },
        lastestPosts: null,
        tags: null
    }),
    methods: {
        getPost,
        getLastestPosts,
        getAllTags
    },
    mounted: function () {
        var idPost = parseInt(this.$route.query.id)

        this.getPost(idPost).then(result => {
            if (result != null) {
                this.item.title = result.title
                this.item.content = result.content
                this.item.image = result.image
            }
        });

        this.getAllTags().then(result => {
            if (result != null) {
                this.tags = result
            }
        });

        this.getLastestPosts(4).then(result => {
            if (result != null) {
                this.lastestPosts = result
            }
        });
    },
};
</script>

<style>
</style>
