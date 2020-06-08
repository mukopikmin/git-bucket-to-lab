<template>
  <div>
    <div class="row">
      <div class="col">
        <div class="form-group">
          <label>GitBucket personal access token</label>
          <input
            v-model="gitbucketToken"
            type="text"
            class="form-control"
            required
          />
        </div>
      </div>
      <div class="col">
        <a
          class="btn btn-outline-primary"
          target="_blank"
          href="http://localhost:8080/root/_application"
        >
          GitBucket personal access token
        </a>
      </div>
    </div>

    <div class="row">
      <div class="col">
        <div class="form-group">
          <label>GitLab personal access token</label>
          <input
            v-model="gitlabToken"
            type="text"
            class="form-control"
            required
          />
        </div>
      </div>
      <div class="col">
        <a
          class="btn btn-outline-primary"
          target="_blank"
          href="http://localhost/profile/personal_access_tokens"
        >
          GitLab personal access token
        </a>
      </div>
    </div>

    <button class="btn btn-outline-primary" @click="auth">Submit</button>
  </div>
</template>

<script>
export default {
  data() {
    return {
      gitbucketToken: '',
      gitlabToken: ''
    }
  },
  methods: {
    auth() {
      localStorage.setItem('gitbucketToken', this.gitbucketToken)
      localStorage.setItem('gitlabToken', this.gitlabToken)

      this.$router.push('/')
    }
  }
}
</script>

<style></style>
