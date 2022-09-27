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
                    <v-card-title>Edite Post</v-card-title>
                    <br>
                    <v-card-subtitle>
                        <form action="" class="form">
                            <v-text-field label="Titre" v-model="title" class="form__control" hide-details="auto" outlined>
                            </v-text-field>
                            <v-select class="mt-3" label="Catégorie" v-model="tagsPost" :items="allTagsName" chips multiple solo outlined></v-select>
                            <v-file-input accept="image/*" label="Image" outlined></v-file-input>
                            <vue-editor v-model="content" />

                            <v-btn color="error" @click="saveBtn" class="login__btn">Sauvegarder</v-btn>
                        </form>
                    </v-card-subtitle>
                </v-card>
            </v-col>
            <v-col xl="12" lg="12" md="12" sm="12" xs="12">
                <PostDetailVue :category="category" :title="title" :content="content" :image="image" />
            </v-col>
        </v-row>
    </v-container>
</div>
</template>

<script>
import SidebarVue from "/components/admin/Sidebar.vue";
import PostDetailVue from "/components/front/PostDetail.vue";

import {
    addPost,
    updatePost,
    getPost,
    getPostTags,
    updatePostTags
} from '/functions/post.js'

import {
    getAllTags
} from '/functions/tag.js'

export default {
    name: "editPostPage",
    head() {
        return {
            title: 'EditPost',
        }
    },
    components: {
        SidebarVue,
        PostDetailVue,
    },
    data: () => ({
        id: 0,
        title: null,
        content: null,
        category: null,
        image: "loading.png",
        tagsPost: [],
        allTags: [],
        allTagsName: [],
    }),
    mounted: function () {
        this.id = parseInt(this.$route.query.id)

        if (this.id != 0) {
            getPost(this.id).then(result => {
                if (result != null) {
                    this.title = result.title
                    this.content = result.content
                    this.image = result.image
                }
            });

            getPostTags(this.id).then(result => {
                if (result != null) {
                    for (const item in result) {
                        this.tagsPost.push(result[item].name)
                    }
                }
            });
        }

        getAllTags().then(result => {
            if (result != null) {
                this.allTags = result
                for (const item in result) {
                    this.allTagsName.push(result[item].name)
                }
            }
        });
    },
    methods: {
        saveBtn() {
            if (this.id != 0) {
                addPost(this.$auth.strategy.token.get(), this.title, this.image, this.content)
            } else {
                updatePost(this.$auth.strategy.token.get(), this.id, this.title, this.image, this.content)
            }

            var tagsId = []
            for (const i in this.tagsPost) {
                for (const j in this.allTags) {
                    if (this.tagsPost[i] === this.allTags[j].name) {
                        tagsId.push(this.allTags[j].id)
                    }
                }
            }

            updatePostTags(this.$auth.strategy.token.get(), this.id, tagsId)

            this.$router.push({
                path: '/admin/post'
            })
        }
    }
};
</script>

<style>
</style>
