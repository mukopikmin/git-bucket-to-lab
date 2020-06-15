<template>
  <div>
    <AuthForm
      :loading="loading"
      :gitbucket-url="gitbucketUrl"
      :gitlab-url="gitlabUrl"
    />
  </div>
</template>

<script>
import AuthForm from '@/components/auth_form'

export default {
  components: {
    AuthForm
  },
  data() {
    return {
      gitbucketUrl: '',
      gitlabUrl: '',
      loading: true
    }
  },
  async mounted() {
    const res = await this.$axios.$get('/auth')

    this.loading = false
    this.gitbucketUrl = res.gitbucket_url
    this.gitlabUrl = res.gitlab_url
  }
}
</script>

<style></style>
