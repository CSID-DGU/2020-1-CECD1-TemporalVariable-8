<template>
  <v-container>
    <v-row>
      <v-col cols="2"></v-col>
      <v-col cols="8">
        <v-form
            ref="form"
            v-model="valid"
            lazy-validation
        >
          <v-text-field
              label="가게명"
              v-model="name"
              :counter="128"
              :rules="nameRules"
              required
          ></v-text-field>

          <v-textarea
              label="가게 설명"
              v-model="description"
              :counter="1024"
              :rules="descriptionRules"
          ></v-textarea>

          <v-select
              label="본점"
              clearable
              v-model="select"
              :items="availableParents"
          ></v-select>

          <v-file-input label="이미지" clearable></v-file-input>
          <v-btn
              color="error"
              class="mr-4"
              @click="reset"
          >
            전부 지우기
          </v-btn>

          <v-btn
              :disabled="!valid"
              color="success"
              @click="send"
          >
            생성하기
          </v-btn>
        </v-form>
      </v-col>
      <v-col cols="2"></v-col>
    </v-row>
  </v-container>
</template>

<script>
export default {
  name: "Add",
  methods: {
    validate() {
      return this.$refs.form.validate()
    },
    reset() {
      this.$refs.form.reset()
    },
    send() {
      let isValid = this.validate()
      if(isValid){
        alert("성공")
      }
    }
  },
  data: () => ({
    valid: true,
    name: '',
    nameRules: [
      v => !!v || '이름은 반드시 입력해야 합니다.',
      v => (v && v.length <= 128) || '이름은 최대 128자 까지입니다.',
    ],
    description: '',
    descriptionRules: [
      v => v ? v.length <= 1024 || '설명은 최대 1024자 까지입니다.': true,
    ],
    select: null,
    availableParents: [
      '동국반점',
    ],
    checkbox: false,
  })
}
</script>

<style scoped>

</style>