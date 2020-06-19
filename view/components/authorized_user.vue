<template>
  <b-card header="Authorization status">
    <div v-if="loading" class="text-center">
      <b-spinner variant="primary"></b-spinner>
    </div>
    <div v-else>
      <b-card-text v-if="isBothAuthorized">
        <p>You are signed in both services succesfully.</p>
        <span> GitBucket: {{ gitbucketUser.login }} </span>
        <b-icon-check-circle-fill variant="success" class="ml-1" />
        <span class="ml-3"> GitLab: {{ gitlabUser.username }} </span>
        <b-icon-check-circle-fill variant="success" class="ml-1" />
        <b-button class="auth" variant="outline-primary" @click="auth"
          >Authorization setting</b-button
        >
      </b-card-text>
    </div>
  </b-card>
</template>

<script>
import { mapState } from 'vuex'
export default {
  props: ['loading'],
  computed: {
    ...mapState(['gitbucketUser', 'gitlabUser']),
    isBothAuthorized() {
      return this.gitbucketUser && this.gitlabUser
    }
  },
  methods: {
    auth() {
      this.$router.push('/auth')
    }
  }
}
</script>

<style scoped>
.auth {
  float: right;
}
</style>
