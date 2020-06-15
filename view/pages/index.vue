<template>
  <div>
    <table class="table table-bordered table-hover">
      <thead>
        <tr>
          <th></th>
          <th>GitBucket</th>
          <th>GitLab</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="(pair, i) in repoProject" :key="i">
          <td>
            <nuxt-link
              :to="{
                name: 'owner-name',
                params: { owner: pair.Repo.owner.login, name: pair.Repo.name }
              }"
              >Migration</nuxt-link
            >
          </td>
          <td>
            <a target="_blank" :href="pair.Repo.http_url">{{
              pair.Repo.full_name
            }}</a>
          </td>
          <td v-if="pair.Project">
            <a target="_blank" :href="pair.Project.web_url">{{
              pair.Project.path_with_namespace
            }}</a>
          </td>
          <td v-else></td>
        </tr>
      </tbody>
    </table>
  </div>
</template>

<script>
import { mapState } from 'vuex'

export default {
  data() {
    return {
      repoProject: []
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

      this.repoProject = await this.getPairs()
    }, 0)
  },
  methods: {
    async getPairs() {
      const res = await this.$axios.$get('', {
        headers: {
          'X-GITBUCKET-TOKEN': this.gitbucketToken,
          'X-GITLAB-TOKEN': this.gitlabToken
        }
      })

      return res.RepoProject
    }
  }
}
</script>

<style></style>
