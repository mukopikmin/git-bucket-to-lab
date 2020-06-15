<template>
  <div>
    <AuhorizedUser
      :loading="loading"
      :gitbucket-user="gitbucketUser"
      :gitlab-user="gitlabUser"
    />

    <div class="mt-3">
      <RepoTable :loading="loading" :pairs="pairs" />
    </div>
  </div>
</template>

<script>
import { mapState } from 'vuex'
import AuhorizedUser from '@/components/authorized_user'
import RepoTable from '@/components/repo_table'

export default {
  components: {
    AuhorizedUser,
    RepoTable
  },
  data() {
    return {
      gitbucketUser: null,
      gitlabUser: null,
      pairs: [],
      loading: true
    }
  },
  computed: {
    ...mapState(['gitbucketToken', 'gitlabToken']),
    isAuthoirized() {
      return !this.gitbucketToken || !this.gitlabToken
    }
  },
  mounted() {
    setTimeout(async () => {
      if (this.isAuthoirized) {
        this.$router.push('/auth')
      }

      const res = await this.$axios.$get('', {
        headers: {
          'X-GITBUCKET-TOKEN': this.gitbucketToken,
          'X-GITLAB-TOKEN': this.gitlabToken
        }
      })
      this.loading = false

      this.gitbucketUser = res.gitbucket_user
      this.gitlabUser = res.gitlab_user
      this.pairs = res.pairs
    }, 0)
  }
}
</script>

<style></style>
