<template>
  <v-card
      class="mx-auto my-12"
      elevation="2">
    <v-img
        height="250"
        position="50% 50%"
        v-if="hasAvatar"
        :src="filePathAvatar"
    ></v-img>
    <v-card-title>{{ formatName }}</v-card-title>
    <v-card-subtitle v-show="hasSubname">{{ formatSubname }}</v-card-subtitle>
    <v-divider class="mx-4"></v-divider>
    <v-card-text>{{ description }}</v-card-text>
    <v-divider class="mx-4"></v-divider>
    <v-card-actions>
      <v-btn
          color="deep-purple lighten-2"
          text
          :to="{path:'/menu', query : { id : rest_name }}"
      >
        메뉴 수정
      </v-btn>
      <v-btn
          color="deep-purple lighten-2"
          text
          :to="{path:'/menu', query : { id : rest_name }}"
      >
        옵션 수정
      </v-btn>
      <v-btn
          color="deep-purple lighten-2"
          text
          :to="{path:'/menu', query : { id : rest_name }}"
      >
        카테고리 수정
      </v-btn>
    </v-card-actions>
    <v-card-actions>
      <v-spacer/>
      <v-btn
          color="deep-purple lighten-2"
          text
          :to="{path:'/order', query : { id : rest_name }}"
      >
        주문 보기
      </v-btn>
    </v-card-actions>
  </v-card>
</template>

<script>
export default {
  name: 'Restaurant',
  mounted() {
  },
  computed: {
    hasSubname() {
      if (this.rest_name.length >= 2) {
        return true
      } else {
        return false
      }
    },
    formatSubname() {
      if (this.hasSubname) {
        return this.rest_name[1]
      }
      return ""
    },
    formatName() {
      if (this.rest_name.length < 1) {
        return "<이름없음>"
      }
      return this.rest_name[0]
    },
    hasAvatar() {
      return typeof this.rest_avatar === "string"
    },
    filePathAvatar() {
      return `/files/${this.rest_avatar}`
    },
    description() {
      if (this.rest_description !== undefined) {
        return this.rest_description
      }
      return ""
    }
  },
  props: {
    rest_id: Number,
    rest_name: Array,
    rest_avatar: String,
    rest_description: String,
  },
  data: () => ({}),
}
</script>
