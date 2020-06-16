<template>
  <div>
    <b-row class="mb-3">
      <b-col sm="6">
        <h4 class="text-center">GitBucket</h4>
        <Repo :repo="repo" :loading="loading" />
      </b-col>
      <b-col sm="6">
        <h4 class="text-center">GitLab</h4>
        <Project :project="project" :loading="loading" />
      </b-col>
    </b-row>

    <b-row class="mb-3">
      <b-col sm="6">
        <RepoIssues :repo="repo" :issues="gitbucketIssues" :loading="loading" />
      </b-col>
      <b-col sm="6">
        <ProjectIssues :issues="gitlabIssues" :loading="loading" />
      </b-col>
    </b-row>

    <b-row class="mb-3">
      <b-col sm="6">
        <PullRequests :repo="repo" :pulls="pulls" :loading="loading" />
      </b-col>
      <b-col sm="6">
        <MergeRequests :merges="merges" :loading="loading" />
      </b-col>
    </b-row>
  </div>
</template>

<script>
import { mapState } from 'vuex'
import Repo from '@/components/repo'
import Project from '@/components/project'
import RepoIssues from '@/components/repo_issues'
import ProjectIssues from '@/components/project_issues'
import PullRequests from '@/components/pull_requests'
import MergeRequests from '@/components/merge_requests'

export default {
  components: {
    Repo,
    Project,
    RepoIssues,
    ProjectIssues,
    PullRequests,
    MergeRequests
  },
  data() {
    return {
      repo: null,
      project: null,
      loading: true
    }
  },
  computed: {
    ...mapState(['gitbucketToken', 'gitlabToken']),
    isAuthoirized() {
      return !this.gitbucketToken || !this.gitlabToken
    },
    gitbucketIssues() {
      return this.repo ? this.repo.issues : []
    },
    gitlabIssues() {
      return this.project ? this.project.issues : []
    },
    pulls() {
      return this.repp ? this.repo.pulls : []
    },
    merges() {
      return this.project ? this.project.merges : []
    }
  },
  mounted() {
    const { owner, name } = this.$nuxt.$route.params

    setTimeout(async () => {
      if (this.isAuthoirized) {
        this.$router.push('/auth')
      }

      const res = await this.$axios.$get(`/${owner}/${name}`, {
        headers: {
          'X-GITBUCKET-TOKEN': this.gitbucketToken,
          'X-GITLAB-TOKEN': this.gitlabToken
        }
      })

      this.repo = res.repo
      this.project = res.project
      this.loading = false
    })
  }
}
</script>

<style></style>
