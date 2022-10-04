<template>
<div class="my-10">
    <v-container>
        <h1 class="m-3">Map</h1>
        <div id="map-wrap" style="height: 90vh">
            <client-only>
                <l-map :zoom=8 :center="[23.791474915526848,90.41672529214094]">
                    <l-tile-layer url="https://{s}.basemaps.cartocdn.com/rastertiles/voyager/{z}/{x}/{y}.png"></l-tile-layer>
                    <l-marker v-for="(location, index) in locations" :key="index" :lat-lng="[location.latitude,location.longitude]">
                        <l-popup>
                            <h4 class="mt-2">{{location.name}}</h4>
                            <v-container ma-0 pa-0 v-for="(postLocation, indexPL) in postLocations" :key="indexPL">
                                <v-container ma-0 pa-0 v-if="postLocation.locationId == location.id">
                                    <v-container ma-0 pa-0 v-for="(post, indexP) in posts" :key="indexP">
                                        <v-container ma-0 pa-0 v-if="postLocation.postId == post.id">
                                            <nuxt-link :to="{ path: '/postDetail?id=' + post.id }">
                                                <h4 class="mt-1">{{post.title}}</h4>
                                            </nuxt-link>
                                        </v-container>
                                    </v-container>
                                </v-container>
                            </v-container>
                        </l-popup>
                        <l-icon>
                            <img height="20px" :src="'http://localhost:8800' + location.image" />
                        </l-icon>
                    </l-marker>
                </l-map>
            </client-only>
        </div>
    </v-container>
</div>
</template>

<script>
import {
    getAllLocations,
} from '/functions/location.js'

import {
    getAllPostLocations,
    getAllPosts
} from '/functions/post.js'

export default {
    name: "MapVue",
    data() {
        return {
            locations: null,
            posts: null,
            postLocations: null,
        }
    },
    mounted: function () {
        getAllLocations().then(result => {
            this.locations = result
        });

        getAllPostLocations().then(result => {
            this.postLocations = result
        });

        getAllPosts().then(result => {
            this.posts = result
        });
    },
};
</script>

<style >
</style>
