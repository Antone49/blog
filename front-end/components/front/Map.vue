<template>
<div class="my-10">
    <v-container>
        <h1 class="m-3">Map</h1>
        <div id="map-wrap" style="height: 90vh">
            <client-only>
                <l-map :zoom=15 :center="[23.791474915526848,90.41672529214094]">
                    <l-tile-layer url="https://{s}.basemaps.cartocdn.com/rastertiles/voyager/{z}/{x}/{y}.png"></l-tile-layer>
                    <l-marker v-for="(item, index) in datas" :key="index" :lat-lng="[item.latitude,item.longitude]">
                        <l-popup>
                            <nuxt-link :to="{ path: '/postDetail?id=' + item.id }">
                                <h4 class="mt-2">{{item.name}}</h4>
                            </nuxt-link>
                        </l-popup>
                        <l-icon>
                            <img height="20px" :src="'http://localhost:8800' + item.image" />
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
    getAllLocations
} from '/functions/location.js'

export default {
    name: "MapVue",
    data() {
        return {
            datas: null,
        }
    },
    mounted: function () {
        getAllLocations().then(result => {
            this['datas'] = result
        });
    },
};
</script>

<style >
</style>
