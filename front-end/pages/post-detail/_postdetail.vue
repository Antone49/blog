<template>
<div>
    <NavbarVue />
    <PostDetailHeroVue :title="item.title" />
    <PostDetailVue :category="item.category" :title="item.title" :content="item.content" :image="item.image" :lastestPosts="lastestPosts" :postTags="postTags" />
    <FooterVue />
</div>
</template>

<script>
import FooterVue from "../../components/front/Footer.vue";
import NavbarVue from "../../components/front/Navbar.vue";
import PostDetailVue from "../../components/front/PostDetail.vue";
import PostDetailHeroVue from "../../components/front/PostDetailHero.vue";

import {
    getPost,
    getLastestPosts,
    getAllTags
} from '../../functions/post.js'

export default {
    name: "PostDetail",
    head() {
        return {
            title: this.item.title,
        }
    },
    components: {
        NavbarVue,
        FooterVue,
        PostDetailHeroVue,
        PostDetailVue,
    },
    data: () => ({
        item: {
            title: null,
            content: null,
            category: null,
            image: "loading.png",
        },
        lastestPosts: null,
        postTags: null
    }),
    methods: {
        getPost,
        getLastestPosts,
        getAllTags
    },
    mounted: function () {
        var pattern = "post-detail/"
        var idPost = parseInt(location.pathname.substring(location.pathname.indexOf(pattern) + pattern.length));

        console.log('idPost %d', idPost);

        this.getPost(idPost).then(result => {
            if (result != null) {
                console.log('result.title %s', result.title);
                this['item'].title = result.title
                this['item'].content = result.content
                this['item'].image = result.image
            }
        });

        this.getAllTags().then(result => {
            if (result != null) {
                this['postTags'] = result
            }
        });

        this.getLastestPosts(4).then(result => {
            if (result != null) {
                this['lastestPosts'] = result
            }
        });
    },
};
</script>

<style>
</style>
