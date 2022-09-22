<template>
<div>
    <SidebarVue />
    <v-container>
        <v-row justify="center" class="mt-12">
            <v-col xl="12" lg="12" md="12" sm="12" xs="12">
                <nuxt-link to="/admin/post">
                    <v-btn color="secondary" class="mb-4">Annuler</v-btn>
                </nuxt-link>
                <v-card>
                    <v-card-title>Update Post</v-card-title>
                    <br>
                    <v-card-subtitle>
                        <form action="" class="form">
                            <v-text-field label="Titre" v-model="item.title" class="form__control" hide-details="auto" outlined>
                            </v-text-field>
                            <v-select class="mt-3" label="Catégorie" outlined></v-select>
                            <v-file-input accept="image/*" label="Image" outlined></v-file-input>
                            <vue-editor v-model="item.content" />
    <div>{{ item.content }}</div>

                            <v-btn color="error" class="login__btn">Update</v-btn>
                        </form>
                    </v-card-subtitle>
                </v-card>
            </v-col>
            <v-col xl="12" lg="12" md="12" sm="12" xs="12">
                <PostDetailVue :category="item.category" :title="item.title" :content="item.content" :image="item.image" />
            </v-col>
        </v-row>
    </v-container>
</div>
</template>

<script>
import SidebarVue from "/components/admin/Sidebar.vue";
import PostDetailVue from "/components/front/PostDetail.vue";

import {
    getPost
} from '/functions/post.js'

export default {
    name: "updatePostPage",
    head() {
        return {
            title: 'Update-Post',
        }
    },
    components: {
        SidebarVue,
        PostDetailVue,
    },
    data: () => ({
        item: {
            title: null,
            content: null,
            category: null,
            image: "loading.png",
        },
    }),
    mounted: function () {
        var idPost = parseInt(this.$route.query.id)

        getPost(idPost).then(result => {
            if (result != null) {
                this.item.title = result.title
                this.item.content = result.content
                this.item.image = result.image
            }
        });
    },
};
</script>

<style>
</style>
