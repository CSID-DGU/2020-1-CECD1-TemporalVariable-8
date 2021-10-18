<template>
  <v-container>
    <v-row>{{ $route.query.id.join(" ") }}</v-row>
    <v-row v-for="menu in menus" :key="menu.id">
      <v-col cols="1">{{ menu.id }}</v-col>
      <v-col cols="1">
        <v-text-field dense v-model="menu.name"></v-text-field>
      </v-col>
      <v-col cols="1">
        <v-text-field dense v-model="menu.description"></v-text-field>
      </v-col>
      <v-col cols="1">
        <v-text-field dense v-model.number="menu.priceAmount" type="number"></v-text-field>
      </v-col>
      <v-col cols="2"></v-col>
      <v-col cols="2">
        <v-select
            v-model="menu.options"
            :items="options"
            item-text="name"
            item-value="id"
            chips
            label="Chips"
            multiple
            solo
        ></v-select>
      </v-col>
      <v-col cols="4">
        <span class="d-flex justify-end">
          <v-btn color="error" @click="real_d(menu)">삭제</v-btn>
        </span>
      </v-col>
    </v-row>
    <v-divider/>
    <v-row>
      <v-col cols="1"></v-col>
      <v-col cols="1">
        <v-text-field dense v-model="addForm.name"></v-text-field>
      </v-col>
      <v-col cols="1">
        <v-text-field dense v-model="addForm.desc"></v-text-field>
      </v-col>
      <v-col cols="1">
        <v-text-field dense v-model.number="addForm.amount" type="number"></v-text-field>
      </v-col>
      <v-col cols="2"></v-col>
      <v-col cols="2">
        <v-select
            v-model="addForm.options"
            :items="options"
            item-text="name"
            item-value="id"
            chips
            label="Chips"
            multiple
            solo
        ></v-select>
      </v-col>
      <v-col cols="4">
        <div class="d-flex justify-end">
          <v-btn color="success" @click="real_c">추가</v-btn>
        </div>
      </v-col>
    </v-row>
  </v-container>
</template>

<script>
export default {
  name: 'Menu',
  mounted() {
    // TODO
    this.loadMenus()
    this.loadOptions()
    // this.fakeLoadMenus()
    // this.fakeLoadOptions()
  },
  methods: {
    mimic_d() {
      this.mimic_c_toggle = !this.mimic_c_toggle
      setTimeout(() => {
        alert('메뉴가 삭제되었습니다.')
        this.menus.pop()
      }, 200)
    },
    mimic_c() {
      this.mimic_c_toggle = !this.mimic_c_toggle
      if (this.mimic_c_toggle) {
        setTimeout(() => alert('이름은 반드시 필요합니다.'), 40)
      } else {
        this.reset()
        setTimeout(() => {
          alert('새로운 메뉴가 추가되었습니다.')
          this.menus.push(
              {
                id: 2,
                name: "짬뽕",
                description: "새빨간 짬뽕",
                priceAmount: 5500,
                currency: 'KRW',
                image: undefined,
                options: [2, 3],
              })
        }, 200)
      }
    },
    real_c() {
      this.$axios.post(`/api/restaurants/${this.$route.query.id.join("/")}/menu`, {
        name: this.addForm.name,
        desc: this.addForm.desc,
        amount: this.addForm.amount,
        options: this.addForm.options,
      }).then(res => {
        console.log(res)
        alert("메뉴가 추가되었습니다.")
        this.loadMenus()
        this.reset()
      }).catch(err => {
        console.log(err)
        alert("서버 오류.")
      })
    },
    real_d(m) {
      console.log(m)
      this.$axios.delete(`/api/restaurants/${this.$route.query.id.join("/")}/menu/${m.id}`).then(res => {
        console.log(res)
        alert("메뉴가 삭제되었습니다.")
        this.loadMenus()
      }).catch(err => {
        console.log(err)
        alert("서버 오류.")
      })
    },
    loadMenus() {
      console.log(`GET : /api/restaurants/${this.$route.query.id.join("/")}/menus`)
      this.$axios.get(`/api/restaurants/${this.$route.query.id.join("/")}/menus`)
          .then(res => {
            this.menus = res.data.menus
          })
          .catch(err => {
            console.log(err)
            this.menus = []
          })
    },
    loadOptions() {
      console.log(`GET : /api/restaurants/${this.$route.query.id.join("/")}/options`)
      this.$axios.get(`/api/restaurants/${this.$route.query.id.join("/")}/options`)
          .then(res => {
            this.options = res.data.menus
          })
          .catch(err => {
            console.log(err)
            this.options = []
          })
    },
    fakeLoadMenus() {
      this.menus = [
        {
          id: 1,
          name: "짜장면",
          description: "까만 짜장면",
          priceAmount: 5000,
          currency: 'KRW',
          image: undefined,
          options: [2],
        },
      ]
    },
    fakeLoadOptions() {
      this.options = [
        {
          id: 2,
          name: "곱빼기",
          priceAmount: 500,
          currency: 'KRW'
        },
        {
          id: 3,
          name: "초곱빼기",
          priceAmount: 1000,
          currency: 'KRW'
        }
      ]
    },
    reset() {
      this.addForm.name = ""
      this.addForm.desc = ""
      this.addForm.amount = ""
      this.addForm.options = []
    }
  },
  data: () => ({
    menus: [],
    options: [],
    addForm: {
      name: "",
      desc: "",
      amount: "",
      currency: "KRW",
      options: [],
    },
  })
}
</script>
