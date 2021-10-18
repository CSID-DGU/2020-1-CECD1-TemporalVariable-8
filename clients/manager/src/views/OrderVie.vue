<template>
  <v-container>
    <v-row>
      <v-col cols="12">
        <div class="d-flex justify-end">
          <v-btn @click="loadOrders()">새로고침</v-btn>
        </div>
      </v-col>
    </v-row>
    <v-row>
      <v-col cols="1">
        <v-card-title>주문번호</v-card-title>
      </v-col>
      <v-col cols="4">
        <v-card-title>주문정보</v-card-title>
      </v-col>
      <v-col cols="1">가격</v-col>
      <v-col cols="4"></v-col>
    </v-row>
    <v-divider/>
    <v-row v-for="or in orders" :key="or.id">
      <v-col cols="1">{{ or.id }}</v-col>
      <v-col cols="4">{{ formatMenus(or.purchased) }}</v-col>
      <v-col cols="1">{{ totalPrice(or.purchased) }} 원</v-col>
      <v-col cols="4">
        <div class="d-flex justify-end">
          <v-btn class="mr-2" color="error" :disabled="or.cooking_at !== undefined" @click="hcooking(or)">조리시작</v-btn>
          <v-btn class="mr-2" color="warning" :disabled="or.cooking_at !== undefined && or.deliver_at !== undefined" @click="hdeliver(or)">배달시작</v-btn>
          <v-btn class="mr-2" color="success" :disabled="or.cooking_at !== undefined && or.deliver_at !== undefined && or.complete_at !== undefined" @click="hcomplete(or)">주문완료
          </v-btn>
        </div>
      </v-col>
    </v-row>

  </v-container>
</template>

<script>
export default {
  name: "OrderVie",
  mounted() {
    this.loadOrders()
  },
  methods: {
    hcooking(target) {
      this.$axios.post(`/api/orders/${target.id}/cooking`)
          .then(() => {
            this.loadOrders()
          })
          .catch(err => {
            console.log(err)
          })
    },
    hdeliver(target) {
      this.$axios.post(`/api/orders/${target.id}/deliver`)
          .then(() => {
            this.loadOrders()
          })
          .catch(err => {
            console.log(err)
          })
    },
    hcomplete(target) {
      this.$axios.post(`/api/orders/${target.id}/complete`)
          .then(() =>{
            this.loadOrders()
          })
          .catch(err => {
            console.log(err)
          })
    },
    formatMenus(menus) {
      let result = []
      for (let m of menus) {
        result.push(`${m.name}x${m.count}`)
      }
      return result.join(", ")
    },
    totalPrice(menus) {
      let result = 0
      for (const m of menus) {
        result += m.amount * m.count
      }
      return result
    },
    loadOrders() {
      this.$axios.get(`/api/restaurants/${this.$route.query.id.join("/")}/orders`)
          .then(res => {
            this.orders = res.data
          })
          .catch(err => {
            console.log(err)
          })
    }
  },
  data: () => ({
    orders: []
  })
}
</script>

<style scoped>

</style>