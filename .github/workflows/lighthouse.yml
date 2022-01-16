name: Test Lighthouse Check
on: [pull_request]

jobs:
  lighthouse-check:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@master
      - run: mkdir ${{ github.workspace }}/tmp/artifacts
      - name: Run Lighthouse
        uses: foo-software/lighthouse-check-action@master
        with:
          awsAccessKeyId: ${{ secrets.LIGHTHOUSE_CHECK_AWS_ACCESS_KEY_ID }}
          awsBucket: ${{ secrets.LIGHTHOUSE_CHECK_AWS_BUCKET }}
          awsRegion: ${{ secrets.LIGHTHOUSE_CHECK_AWS_REGION }}
          awsSecretAccessKey: ${{ secrets.LIGHTHOUSE_CHECK_AWS_SECRET_ACCESS_KEY }}
          gitAuthor: ${{ github.actor }}
          gitBranch: ${{ github.ref }}
          gitHubAccessToken: ${{ secrets.LIGHTHOUSE_CHECK_GITHUB_ACCESS_TOKEN }}
          outputDirectory: ${{ github.workspace }}/tmp/artifacts
          urls: 'https://www.foo.software,https://www.foo.software/contact'
          sha: ${{ github.sha }}
          slackWebhookUrl: ${{ secrets.LIGHTHOUSE_CHECK_WEBHOOK_URL }}
