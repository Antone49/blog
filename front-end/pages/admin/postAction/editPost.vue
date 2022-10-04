<template>
<div>
    <SidebarVue />
    <v-container>
        <v-row justify="center" class="mt-12">
            <v-col xl="12" lg="12" md="12" sm="12" xs="12">
                <nuxt-link to="/admin/post">
                    <v-btn color="secondary" @click="cancelBtn" class="mb-4">Annuler</v-btn>
                </nuxt-link>
                <v-card>
                    <v-card-title>Edite Post</v-card-title>
                    <br>
                    <v-card-subtitle>
                        <form action="" class="form">
                            <v-text-field label="Titre" v-model="title" class="form__control" hide-details="auto" outlined>
                            </v-text-field>

                            <v-select class="mt-3" label="Catégories" v-model="tagsPost" :items="allTags" multiple outlined>
                                <template v-slot:selection="{ item }">
                                    <v-chip>
                                        {{item.name}}
                                    </v-chip>
                                </template>
                                <template v-slot:item="{item}">
                                    <v-icon v-if="tagsPost !== null && tagsPost.includes(item)" color="primary" class="mr-3">
                                        mdi-checkbox-marked
                                    </v-icon>
                                    <v-icon v-else class="mr-3">
                                        mdi-checkbox-blank-outline
                                    </v-icon>
                                    {{item.name}}
                                </template>
                            </v-select>

                            <v-select class="mt-3" label="Locations" v-model="locationsPost" :items="allLocations" multiple outlined>
                                <template v-slot:selection="{ item, index }">
                                    <v-chip>
                                        <img :src="'http://localhost:8800' + item.image" width="25" height="25">
                                        {{item.name}}
                                    </v-chip>
                                </template>
                                <template v-slot:item="{item}">
                                    <v-icon v-if="locationsPost !== null && locationsPost.includes(item)" color="primary" class="mr-3">
                                        mdi-checkbox-marked
                                    </v-icon>
                                    <v-icon v-else class="mr-3">
                                        mdi-checkbox-blank-outline
                                    </v-icon>
                                    <img :src="'http://localhost:8800' + item.image" width="30" height="30">
                                    {{item.name}}
                                </template>
                            </v-select>

                            <v-file-input accept="image/*" label="Image" chips outlined v-model="file" @change="previewImage"></v-file-input>

                            <textarea v-model="content" id="file-picker"></textarea>
                            <v-btn color="error" @click="saveBtn" class="login__btn">Sauvegarder</v-btn>
                        </form>
                    </v-card-subtitle>
                </v-card>
            </v-col>
            <v-col xl="12" lg="12" md="12" sm="12" xs="12">
                <PostDetailVue :category="category" :title="title" :content="content" :image="imageView" />
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
    updatePostTags,
    getPostLocations,
    updatePostLocations,
    updatePostImage,
    removePost
} from '/functions/post.js'

import {
    getAllLocations
} from '/functions/location.js'

import {
    getAllTags
} from '/functions/tag.js'

import {
    uploadImage
} from '/functions/image.js'

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
        id: NaN,
        isNewPost: false,
        title: null,
        content: null,
        category: null,
        image: null,
        imageView: "http://localhost:8800/images/loading.png",
        allTags: [],
        tagsPost: [],
        allLocations: [],
        locationsPost: [],
        file: null,
        readers: null,
    }),
    mounted: function () {
        this.id = parseInt(this.$route.query.id)

        getAllTags().then(result => {
            if (result != null) {
                this.allTags = result

                if (!isNaN(this.id)) {
                    getPostTags(this.id).then(result => {
                        if (result != null) {
                            for (const i in result) {
                                for (const j in this.allTags) {
                                    if (result[i].id == this.allTags[j].id) {
                                        this.tagsPost.push(this.allTags[j])
                                    }
                                }
                            }
                        }
                    });
                }
            }
        });

        getAllLocations().then(result => {
            if (result != null) {
                this.allLocations = result

                if (!isNaN(this.id)) {
                    getPostLocations(this.id).then(result => {
                        if (result != null) {
                            for (const i in result) {
                                for (const j in this.allLocations) {
                                    if (result[i].id == this.allLocations[j].id) {
                                        this.locationsPost.push(this.allLocations[j])
                                    }
                                }
                            }
                        }
                    });
                }
            }
        });

        if (!isNaN(this.id)) {
            getPost(this.id).then(result => {
                if (result != null) {
                    this.title = result.title
                    this.content = result.content
                    this.image = result.image
                    this.imageView = "http://localhost:8800" + this.image
                }
            });
        } else {
            addPost(this.$auth.strategy.token.get()).then(result => {
                if (result != null) {
                    this.id = result.id
                    this.isNewPost = true
                }
            });
        }

        tinymce.init({
            selector: 'textarea#file-picker',
            plugins: 'lists link image table code help wordcount',
            toolbar: 'undo redo | styleselect | bold italic | alignleft aligncenter alignright alignjustify | bullist numlist outdent indent | link image table | fontsizeselect | code',
            skin: "oxide-dark",
            content_css: "dark",
            /* enable title field in the Image dialog*/
            image_title: true,
            file_picker_types: 'image',
            /* and here's our custom image picker*/
            file_picker_callback: (cb, value, meta) => {
                const input = document.createElement('input');
                input.setAttribute('type', 'file');
                input.setAttribute('accept', 'image/*');

                input.addEventListener('change', (e) => {
                    const file = e.target.files[0];

                    const reader = new FileReader();
                    reader.addEventListener('load', () => {
                        uploadImage(this.$auth.strategy.token.get(), file).then(result => {
                            cb("http://localhost:8800" + result.filename, {
                                title: file.name
                            });
                        });
                    });
                    reader.readAsDataURL(file);
                });

                input.click();
            },
            image_upload_handler_callback: (blobInfo, progress) => new Promise((resolve, reject) => {
                debugger
                uploadImage(this.$auth.strategy.token.get(), blobInfo.blob()).then(result => {
                   debugger
                });
            }),

            content_style: 'body { font-family:Helvetica,Arial,sans-serif; font-size:16px }'
        });
    },
    methods: {
        saveBtn() {
            var content = tinymce.activeEditor.getContent();

            if (!isNaN(this.id)) {
                updatePost(this.$auth.strategy.token.get(), this.id, this.title, content)

                var tagsId = []
                for (const i in this.tagsPost) {
                    tagsId.push(this.tagsPost[i].id)
                }

                updatePostTags(this.$auth.strategy.token.get(), this.id, tagsId)

                var locationId = []
                for (const i in this.locationsPost) {
                    locationId.push(this.locationsPost[i].id)
                }

                updatePostLocations(this.$auth.strategy.token.get(), this.id, locationId)

                if (this.file) {
                    uploadImage(this.$auth.strategy.token.get(), this.file).then(result => {
                        updatePostImage(this.$auth.strategy.token.get(), this.id, result.filename)
                    });
                }
            }

            this.$router.push({
                path: '/admin/post'
            })
        },
        previewImage() {
            const reader = new FileReader();
            reader.readAsDataURL(this.file);
            reader.onload = e => {
                this.imageView = e.target.result;
            };
        },
        cancelBtn() {
            if (this.isNewPost) {
                removePost(this.$auth.strategy.token.get(), this.id)
            }
        },
    }
};
</script>

<style>
</style>
