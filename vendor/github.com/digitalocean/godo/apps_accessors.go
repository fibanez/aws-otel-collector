// Code generated automatically. DO NOT EDIT.

package godo

import (
	"time"
)

// GetActiveDeployment returns the ActiveDeployment field.
func (a *App) GetActiveDeployment() *Deployment {
	if a == nil {
		return nil
	}
	return a.ActiveDeployment
}

// GetBuildConfig returns the BuildConfig field.
func (a *App) GetBuildConfig() *AppBuildConfig {
	if a == nil {
		return nil
	}
	return a.BuildConfig
}

// GetCreatedAt returns the CreatedAt field.
func (a *App) GetCreatedAt() time.Time {
	if a == nil {
		return time.Time{}
	}
	return a.CreatedAt
}

// GetDefaultIngress returns the DefaultIngress field.
func (a *App) GetDefaultIngress() string {
	if a == nil {
		return ""
	}
	return a.DefaultIngress
}

// GetDomains returns the Domains field.
func (a *App) GetDomains() []*AppDomain {
	if a == nil {
		return nil
	}
	return a.Domains
}

// GetID returns the ID field.
func (a *App) GetID() string {
	if a == nil {
		return ""
	}
	return a.ID
}

// GetInProgressDeployment returns the InProgressDeployment field.
func (a *App) GetInProgressDeployment() *Deployment {
	if a == nil {
		return nil
	}
	return a.InProgressDeployment
}

// GetLastDeploymentActiveAt returns the LastDeploymentActiveAt field.
func (a *App) GetLastDeploymentActiveAt() time.Time {
	if a == nil {
		return time.Time{}
	}
	return a.LastDeploymentActiveAt
}

// GetLastDeploymentCreatedAt returns the LastDeploymentCreatedAt field.
func (a *App) GetLastDeploymentCreatedAt() time.Time {
	if a == nil {
		return time.Time{}
	}
	return a.LastDeploymentCreatedAt
}

// GetLiveDomain returns the LiveDomain field.
func (a *App) GetLiveDomain() string {
	if a == nil {
		return ""
	}
	return a.LiveDomain
}

// GetLiveURL returns the LiveURL field.
func (a *App) GetLiveURL() string {
	if a == nil {
		return ""
	}
	return a.LiveURL
}

// GetLiveURLBase returns the LiveURLBase field.
func (a *App) GetLiveURLBase() string {
	if a == nil {
		return ""
	}
	return a.LiveURLBase
}

// GetOwnerUUID returns the OwnerUUID field.
func (a *App) GetOwnerUUID() string {
	if a == nil {
		return ""
	}
	return a.OwnerUUID
}

// GetPendingDeployment returns the PendingDeployment field.
func (a *App) GetPendingDeployment() *Deployment {
	if a == nil {
		return nil
	}
	return a.PendingDeployment
}

// GetPinnedDeployment returns the PinnedDeployment field.
func (a *App) GetPinnedDeployment() *Deployment {
	if a == nil {
		return nil
	}
	return a.PinnedDeployment
}

// GetProjectID returns the ProjectID field.
func (a *App) GetProjectID() string {
	if a == nil {
		return ""
	}
	return a.ProjectID
}

// GetRegion returns the Region field.
func (a *App) GetRegion() *AppRegion {
	if a == nil {
		return nil
	}
	return a.Region
}

// GetSpec returns the Spec field.
func (a *App) GetSpec() *AppSpec {
	if a == nil {
		return nil
	}
	return a.Spec
}

// GetTierSlug returns the TierSlug field.
func (a *App) GetTierSlug() string {
	if a == nil {
		return ""
	}
	return a.TierSlug
}

// GetUpdatedAt returns the UpdatedAt field.
func (a *App) GetUpdatedAt() time.Time {
	if a == nil {
		return time.Time{}
	}
	return a.UpdatedAt
}

// GetComponentName returns the ComponentName field.
func (a *AppAlert) GetComponentName() string {
	if a == nil {
		return ""
	}
	return a.ComponentName
}

// GetEmails returns the Emails field.
func (a *AppAlert) GetEmails() []string {
	if a == nil {
		return nil
	}
	return a.Emails
}

// GetID returns the ID field.
func (a *AppAlert) GetID() string {
	if a == nil {
		return ""
	}
	return a.ID
}

// GetPhase returns the Phase field.
func (a *AppAlert) GetPhase() AppAlertPhase {
	if a == nil {
		return ""
	}
	return a.Phase
}

// GetProgress returns the Progress field.
func (a *AppAlert) GetProgress() *AppAlertProgress {
	if a == nil {
		return nil
	}
	return a.Progress
}

// GetSlackWebhooks returns the SlackWebhooks field.
func (a *AppAlert) GetSlackWebhooks() []*AppAlertSlackWebhook {
	if a == nil {
		return nil
	}
	return a.SlackWebhooks
}

// GetSpec returns the Spec field.
func (a *AppAlert) GetSpec() *AppAlertSpec {
	if a == nil {
		return nil
	}
	return a.Spec
}

// GetSteps returns the Steps field.
func (a *AppAlertProgress) GetSteps() []*AppAlertProgressStep {
	if a == nil {
		return nil
	}
	return a.Steps
}

// GetEndedAt returns the EndedAt field.
func (a *AppAlertProgressStep) GetEndedAt() time.Time {
	if a == nil {
		return time.Time{}
	}
	return a.EndedAt
}

// GetName returns the Name field.
func (a *AppAlertProgressStep) GetName() string {
	if a == nil {
		return ""
	}
	return a.Name
}

// GetReason returns the Reason field.
func (a *AppAlertProgressStep) GetReason() *AppAlertProgressStepReason {
	if a == nil {
		return nil
	}
	return a.Reason
}

// GetStartedAt returns the StartedAt field.
func (a *AppAlertProgressStep) GetStartedAt() time.Time {
	if a == nil {
		return time.Time{}
	}
	return a.StartedAt
}

// GetStatus returns the Status field.
func (a *AppAlertProgressStep) GetStatus() AppAlertProgressStepStatus {
	if a == nil {
		return ""
	}
	return a.Status
}

// GetSteps returns the Steps field.
func (a *AppAlertProgressStep) GetSteps() []*AppAlertProgressStep {
	if a == nil {
		return nil
	}
	return a.Steps
}

// GetCode returns the Code field.
func (a *AppAlertProgressStepReason) GetCode() string {
	if a == nil {
		return ""
	}
	return a.Code
}

// GetMessage returns the Message field.
func (a *AppAlertProgressStepReason) GetMessage() string {
	if a == nil {
		return ""
	}
	return a.Message
}

// GetChannel returns the Channel field.
func (a *AppAlertSlackWebhook) GetChannel() string {
	if a == nil {
		return ""
	}
	return a.Channel
}

// GetURL returns the URL field.
func (a *AppAlertSlackWebhook) GetURL() string {
	if a == nil {
		return ""
	}
	return a.URL
}

// GetDisabled returns the Disabled field.
func (a *AppAlertSpec) GetDisabled() bool {
	if a == nil {
		return false
	}
	return a.Disabled
}

// GetOperator returns the Operator field.
func (a *AppAlertSpec) GetOperator() AppAlertSpecOperator {
	if a == nil {
		return ""
	}
	return a.Operator
}

// GetRule returns the Rule field.
func (a *AppAlertSpec) GetRule() AppAlertSpecRule {
	if a == nil {
		return ""
	}
	return a.Rule
}

// GetValue returns the Value field.
func (a *AppAlertSpec) GetValue() float32 {
	if a == nil {
		return 0
	}
	return a.Value
}

// GetWindow returns the Window field.
func (a *AppAlertSpec) GetWindow() AppAlertSpecWindow {
	if a == nil {
		return ""
	}
	return a.Window
}

// GetCNBVersioning returns the CNBVersioning field.
func (a *AppBuildConfig) GetCNBVersioning() *AppBuildConfigCNBVersioning {
	if a == nil {
		return nil
	}
	return a.CNBVersioning
}

// GetBuildpacks returns the Buildpacks field.
func (a *AppBuildConfigCNBVersioning) GetBuildpacks() []*Buildpack {
	if a == nil {
		return nil
	}
	return a.Buildpacks
}

// GetStackID returns the StackID field.
func (a *AppBuildConfigCNBVersioning) GetStackID() string {
	if a == nil {
		return ""
	}
	return a.StackID
}

// GetAllowCredentials returns the AllowCredentials field.
func (a *AppCORSPolicy) GetAllowCredentials() bool {
	if a == nil {
		return false
	}
	return a.AllowCredentials
}

// GetAllowHeaders returns the AllowHeaders field.
func (a *AppCORSPolicy) GetAllowHeaders() []string {
	if a == nil {
		return nil
	}
	return a.AllowHeaders
}

// GetAllowMethods returns the AllowMethods field.
func (a *AppCORSPolicy) GetAllowMethods() []string {
	if a == nil {
		return nil
	}
	return a.AllowMethods
}

// GetAllowOrigins returns the AllowOrigins field.
func (a *AppCORSPolicy) GetAllowOrigins() []*AppStringMatch {
	if a == nil {
		return nil
	}
	return a.AllowOrigins
}

// GetExposeHeaders returns the ExposeHeaders field.
func (a *AppCORSPolicy) GetExposeHeaders() []string {
	if a == nil {
		return nil
	}
	return a.ExposeHeaders
}

// GetMaxAge returns the MaxAge field.
func (a *AppCORSPolicy) GetMaxAge() string {
	if a == nil {
		return ""
	}
	return a.MaxAge
}

// GetProjectID returns the ProjectID field.
func (a *AppCreateRequest) GetProjectID() string {
	if a == nil {
		return ""
	}
	return a.ProjectID
}

// GetSpec returns the Spec field.
func (a *AppCreateRequest) GetSpec() *AppSpec {
	if a == nil {
		return nil
	}
	return a.Spec
}

// GetClusterName returns the ClusterName field.
func (a *AppDatabaseSpec) GetClusterName() string {
	if a == nil {
		return ""
	}
	return a.ClusterName
}

// GetDBName returns the DBName field.
func (a *AppDatabaseSpec) GetDBName() string {
	if a == nil {
		return ""
	}
	return a.DBName
}

// GetDBUser returns the DBUser field.
func (a *AppDatabaseSpec) GetDBUser() string {
	if a == nil {
		return ""
	}
	return a.DBUser
}

// GetEngine returns the Engine field.
func (a *AppDatabaseSpec) GetEngine() AppDatabaseSpecEngine {
	if a == nil {
		return ""
	}
	return a.Engine
}

// GetName returns the Name field.
func (a *AppDatabaseSpec) GetName() string {
	if a == nil {
		return ""
	}
	return a.Name
}

// GetNumNodes returns the NumNodes field.
func (a *AppDatabaseSpec) GetNumNodes() int64 {
	if a == nil {
		return 0
	}
	return a.NumNodes
}

// GetProduction returns the Production field.
func (a *AppDatabaseSpec) GetProduction() bool {
	if a == nil {
		return false
	}
	return a.Production
}

// GetSize returns the Size field.
func (a *AppDatabaseSpec) GetSize() string {
	if a == nil {
		return ""
	}
	return a.Size
}

// GetVersion returns the Version field.
func (a *AppDatabaseSpec) GetVersion() string {
	if a == nil {
		return ""
	}
	return a.Version
}

// GetCertificateExpiresAt returns the CertificateExpiresAt field.
func (a *AppDomain) GetCertificateExpiresAt() time.Time {
	if a == nil {
		return time.Time{}
	}
	return a.CertificateExpiresAt
}

// GetID returns the ID field.
func (a *AppDomain) GetID() string {
	if a == nil {
		return ""
	}
	return a.ID
}

// GetPhase returns the Phase field.
func (a *AppDomain) GetPhase() AppDomainPhase {
	if a == nil {
		return ""
	}
	return a.Phase
}

// GetProgress returns the Progress field.
func (a *AppDomain) GetProgress() *AppDomainProgress {
	if a == nil {
		return nil
	}
	return a.Progress
}

// GetRotateValidationRecords returns the RotateValidationRecords field.
func (a *AppDomain) GetRotateValidationRecords() bool {
	if a == nil {
		return false
	}
	return a.RotateValidationRecords
}

// GetSpec returns the Spec field.
func (a *AppDomain) GetSpec() *AppDomainSpec {
	if a == nil {
		return nil
	}
	return a.Spec
}

// GetValidation returns the Validation field.
func (a *AppDomain) GetValidation() *AppDomainValidation {
	if a == nil {
		return nil
	}
	return a.Validation
}

// GetValidations returns the Validations field.
func (a *AppDomain) GetValidations() []*AppDomainValidation {
	if a == nil {
		return nil
	}
	return a.Validations
}

// GetSteps returns the Steps field.
func (a *AppDomainProgress) GetSteps() []*AppDomainProgressStep {
	if a == nil {
		return nil
	}
	return a.Steps
}

// GetEndedAt returns the EndedAt field.
func (a *AppDomainProgressStep) GetEndedAt() time.Time {
	if a == nil {
		return time.Time{}
	}
	return a.EndedAt
}

// GetName returns the Name field.
func (a *AppDomainProgressStep) GetName() string {
	if a == nil {
		return ""
	}
	return a.Name
}

// GetReason returns the Reason field.
func (a *AppDomainProgressStep) GetReason() *AppDomainProgressStepReason {
	if a == nil {
		return nil
	}
	return a.Reason
}

// GetStartedAt returns the StartedAt field.
func (a *AppDomainProgressStep) GetStartedAt() time.Time {
	if a == nil {
		return time.Time{}
	}
	return a.StartedAt
}

// GetStatus returns the Status field.
func (a *AppDomainProgressStep) GetStatus() AppDomainProgressStepStatus {
	if a == nil {
		return ""
	}
	return a.Status
}

// GetSteps returns the Steps field.
func (a *AppDomainProgressStep) GetSteps() []*AppDomainProgressStep {
	if a == nil {
		return nil
	}
	return a.Steps
}

// GetCode returns the Code field.
func (a *AppDomainProgressStepReason) GetCode() string {
	if a == nil {
		return ""
	}
	return a.Code
}

// GetMessage returns the Message field.
func (a *AppDomainProgressStepReason) GetMessage() string {
	if a == nil {
		return ""
	}
	return a.Message
}

// GetCertificate returns the Certificate field.
func (a *AppDomainSpec) GetCertificate() string {
	if a == nil {
		return ""
	}
	return a.Certificate
}

// GetDomain returns the Domain field.
func (a *AppDomainSpec) GetDomain() string {
	if a == nil {
		return ""
	}
	return a.Domain
}

// GetMinimumTLSVersion returns the MinimumTLSVersion field.
func (a *AppDomainSpec) GetMinimumTLSVersion() string {
	if a == nil {
		return ""
	}
	return a.MinimumTLSVersion
}

// GetType returns the Type field.
func (a *AppDomainSpec) GetType() AppDomainSpecType {
	if a == nil {
		return ""
	}
	return a.Type
}

// GetWildcard returns the Wildcard field.
func (a *AppDomainSpec) GetWildcard() bool {
	if a == nil {
		return false
	}
	return a.Wildcard
}

// GetZone returns the Zone field.
func (a *AppDomainSpec) GetZone() string {
	if a == nil {
		return ""
	}
	return a.Zone
}

// GetTXTName returns the TXTName field.
func (a *AppDomainValidation) GetTXTName() string {
	if a == nil {
		return ""
	}
	return a.TXTName
}

// GetTXTValue returns the TXTValue field.
func (a *AppDomainValidation) GetTXTValue() string {
	if a == nil {
		return ""
	}
	return a.TXTValue
}

// GetAlerts returns the Alerts field.
func (a *AppFunctionsSpec) GetAlerts() []*AppAlertSpec {
	if a == nil {
		return nil
	}
	return a.Alerts
}

// GetCORS returns the CORS field.
func (a *AppFunctionsSpec) GetCORS() *AppCORSPolicy {
	if a == nil {
		return nil
	}
	return a.CORS
}

// GetEnvs returns the Envs field.
func (a *AppFunctionsSpec) GetEnvs() []*AppVariableDefinition {
	if a == nil {
		return nil
	}
	return a.Envs
}

// GetGit returns the Git field.
func (a *AppFunctionsSpec) GetGit() *GitSourceSpec {
	if a == nil {
		return nil
	}
	return a.Git
}

// GetGitHub returns the GitHub field.
func (a *AppFunctionsSpec) GetGitHub() *GitHubSourceSpec {
	if a == nil {
		return nil
	}
	return a.GitHub
}

// GetGitLab returns the GitLab field.
func (a *AppFunctionsSpec) GetGitLab() *GitLabSourceSpec {
	if a == nil {
		return nil
	}
	return a.GitLab
}

// GetLogDestinations returns the LogDestinations field.
func (a *AppFunctionsSpec) GetLogDestinations() []*AppLogDestinationSpec {
	if a == nil {
		return nil
	}
	return a.LogDestinations
}

// GetName returns the Name field.
func (a *AppFunctionsSpec) GetName() string {
	if a == nil {
		return ""
	}
	return a.Name
}

// GetRoutes returns the Routes field.
func (a *AppFunctionsSpec) GetRoutes() []*AppRouteSpec {
	if a == nil {
		return nil
	}
	return a.Routes
}

// GetSourceDir returns the SourceDir field.
func (a *AppFunctionsSpec) GetSourceDir() string {
	if a == nil {
		return ""
	}
	return a.SourceDir
}

// GetLoadBalancer returns the LoadBalancer field.
func (a *AppIngressSpec) GetLoadBalancer() AppIngressSpecLoadBalancer {
	if a == nil {
		return ""
	}
	return a.LoadBalancer
}

// GetLoadBalancerSize returns the LoadBalancerSize field.
func (a *AppIngressSpec) GetLoadBalancerSize() int64 {
	if a == nil {
		return 0
	}
	return a.LoadBalancerSize
}

// GetRules returns the Rules field.
func (a *AppIngressSpec) GetRules() []*AppIngressSpecRule {
	if a == nil {
		return nil
	}
	return a.Rules
}

// GetComponent returns the Component field.
func (a *AppIngressSpecRule) GetComponent() *AppIngressSpecRuleRoutingComponent {
	if a == nil {
		return nil
	}
	return a.Component
}

// GetCORS returns the CORS field.
func (a *AppIngressSpecRule) GetCORS() *AppCORSPolicy {
	if a == nil {
		return nil
	}
	return a.CORS
}

// GetMatch returns the Match field.
func (a *AppIngressSpecRule) GetMatch() *AppIngressSpecRuleMatch {
	if a == nil {
		return nil
	}
	return a.Match
}

// GetRedirect returns the Redirect field.
func (a *AppIngressSpecRule) GetRedirect() *AppIngressSpecRuleRoutingRedirect {
	if a == nil {
		return nil
	}
	return a.Redirect
}

// GetPath returns the Path field.
func (a *AppIngressSpecRuleMatch) GetPath() *AppIngressSpecRuleStringMatch {
	if a == nil {
		return nil
	}
	return a.Path
}

// GetName returns the Name field.
func (a *AppIngressSpecRuleRoutingComponent) GetName() string {
	if a == nil {
		return ""
	}
	return a.Name
}

// GetPreservePathPrefix returns the PreservePathPrefix field.
func (a *AppIngressSpecRuleRoutingComponent) GetPreservePathPrefix() bool {
	if a == nil {
		return false
	}
	return a.PreservePathPrefix
}

// GetRewrite returns the Rewrite field.
func (a *AppIngressSpecRuleRoutingComponent) GetRewrite() string {
	if a == nil {
		return ""
	}
	return a.Rewrite
}

// GetAuthority returns the Authority field.
func (a *AppIngressSpecRuleRoutingRedirect) GetAuthority() string {
	if a == nil {
		return ""
	}
	return a.Authority
}

// GetPort returns the Port field.
func (a *AppIngressSpecRuleRoutingRedirect) GetPort() int64 {
	if a == nil {
		return 0
	}
	return a.Port
}

// GetRedirectCode returns the RedirectCode field.
func (a *AppIngressSpecRuleRoutingRedirect) GetRedirectCode() int64 {
	if a == nil {
		return 0
	}
	return a.RedirectCode
}

// GetScheme returns the Scheme field.
func (a *AppIngressSpecRuleRoutingRedirect) GetScheme() string {
	if a == nil {
		return ""
	}
	return a.Scheme
}

// GetUri returns the Uri field.
func (a *AppIngressSpecRuleRoutingRedirect) GetUri() string {
	if a == nil {
		return ""
	}
	return a.Uri
}

// GetPrefix returns the Prefix field.
func (a *AppIngressSpecRuleStringMatch) GetPrefix() string {
	if a == nil {
		return ""
	}
	return a.Prefix
}

// GetCPUs returns the CPUs field.
func (a *AppInstanceSize) GetCPUs() string {
	if a == nil {
		return ""
	}
	return a.CPUs
}

// GetCPUType returns the CPUType field.
func (a *AppInstanceSize) GetCPUType() AppInstanceSizeCPUType {
	if a == nil {
		return ""
	}
	return a.CPUType
}

// GetMemoryBytes returns the MemoryBytes field.
func (a *AppInstanceSize) GetMemoryBytes() string {
	if a == nil {
		return ""
	}
	return a.MemoryBytes
}

// GetName returns the Name field.
func (a *AppInstanceSize) GetName() string {
	if a == nil {
		return ""
	}
	return a.Name
}

// GetSlug returns the Slug field.
func (a *AppInstanceSize) GetSlug() string {
	if a == nil {
		return ""
	}
	return a.Slug
}

// GetTierDowngradeTo returns the TierDowngradeTo field.
func (a *AppInstanceSize) GetTierDowngradeTo() string {
	if a == nil {
		return ""
	}
	return a.TierDowngradeTo
}

// GetTierSlug returns the TierSlug field.
func (a *AppInstanceSize) GetTierSlug() string {
	if a == nil {
		return ""
	}
	return a.TierSlug
}

// GetTierUpgradeTo returns the TierUpgradeTo field.
func (a *AppInstanceSize) GetTierUpgradeTo() string {
	if a == nil {
		return ""
	}
	return a.TierUpgradeTo
}

// GetUSDPerMonth returns the USDPerMonth field.
func (a *AppInstanceSize) GetUSDPerMonth() string {
	if a == nil {
		return ""
	}
	return a.USDPerMonth
}

// GetUSDPerSecond returns the USDPerSecond field.
func (a *AppInstanceSize) GetUSDPerSecond() string {
	if a == nil {
		return ""
	}
	return a.USDPerSecond
}

// GetAlerts returns the Alerts field.
func (a *AppJobSpec) GetAlerts() []*AppAlertSpec {
	if a == nil {
		return nil
	}
	return a.Alerts
}

// GetBuildCommand returns the BuildCommand field.
func (a *AppJobSpec) GetBuildCommand() string {
	if a == nil {
		return ""
	}
	return a.BuildCommand
}

// GetDockerfilePath returns the DockerfilePath field.
func (a *AppJobSpec) GetDockerfilePath() string {
	if a == nil {
		return ""
	}
	return a.DockerfilePath
}

// GetEnvironmentSlug returns the EnvironmentSlug field.
func (a *AppJobSpec) GetEnvironmentSlug() string {
	if a == nil {
		return ""
	}
	return a.EnvironmentSlug
}

// GetEnvs returns the Envs field.
func (a *AppJobSpec) GetEnvs() []*AppVariableDefinition {
	if a == nil {
		return nil
	}
	return a.Envs
}

// GetGit returns the Git field.
func (a *AppJobSpec) GetGit() *GitSourceSpec {
	if a == nil {
		return nil
	}
	return a.Git
}

// GetGitHub returns the GitHub field.
func (a *AppJobSpec) GetGitHub() *GitHubSourceSpec {
	if a == nil {
		return nil
	}
	return a.GitHub
}

// GetGitLab returns the GitLab field.
func (a *AppJobSpec) GetGitLab() *GitLabSourceSpec {
	if a == nil {
		return nil
	}
	return a.GitLab
}

// GetImage returns the Image field.
func (a *AppJobSpec) GetImage() *ImageSourceSpec {
	if a == nil {
		return nil
	}
	return a.Image
}

// GetInstanceCount returns the InstanceCount field.
func (a *AppJobSpec) GetInstanceCount() int64 {
	if a == nil {
		return 0
	}
	return a.InstanceCount
}

// GetInstanceSizeSlug returns the InstanceSizeSlug field.
func (a *AppJobSpec) GetInstanceSizeSlug() string {
	if a == nil {
		return ""
	}
	return a.InstanceSizeSlug
}

// GetKind returns the Kind field.
func (a *AppJobSpec) GetKind() AppJobSpecKind {
	if a == nil {
		return ""
	}
	return a.Kind
}

// GetLogDestinations returns the LogDestinations field.
func (a *AppJobSpec) GetLogDestinations() []*AppLogDestinationSpec {
	if a == nil {
		return nil
	}
	return a.LogDestinations
}

// GetName returns the Name field.
func (a *AppJobSpec) GetName() string {
	if a == nil {
		return ""
	}
	return a.Name
}

// GetRunCommand returns the RunCommand field.
func (a *AppJobSpec) GetRunCommand() string {
	if a == nil {
		return ""
	}
	return a.RunCommand
}

// GetSourceDir returns the SourceDir field.
func (a *AppJobSpec) GetSourceDir() string {
	if a == nil {
		return ""
	}
	return a.SourceDir
}

// GetDatadog returns the Datadog field.
func (a *AppLogDestinationSpec) GetDatadog() *AppLogDestinationSpecDataDog {
	if a == nil {
		return nil
	}
	return a.Datadog
}

// GetEndpoint returns the Endpoint field.
func (a *AppLogDestinationSpec) GetEndpoint() string {
	if a == nil {
		return ""
	}
	return a.Endpoint
}

// GetHeaders returns the Headers field.
func (a *AppLogDestinationSpec) GetHeaders() []*AppLogDestinationSpecHeader {
	if a == nil {
		return nil
	}
	return a.Headers
}

// GetLogtail returns the Logtail field.
func (a *AppLogDestinationSpec) GetLogtail() *AppLogDestinationSpecLogtail {
	if a == nil {
		return nil
	}
	return a.Logtail
}

// GetName returns the Name field.
func (a *AppLogDestinationSpec) GetName() string {
	if a == nil {
		return ""
	}
	return a.Name
}

// GetPapertrail returns the Papertrail field.
func (a *AppLogDestinationSpec) GetPapertrail() *AppLogDestinationSpecPapertrail {
	if a == nil {
		return nil
	}
	return a.Papertrail
}

// GetTLSInsecure returns the TLSInsecure field.
func (a *AppLogDestinationSpec) GetTLSInsecure() bool {
	if a == nil {
		return false
	}
	return a.TLSInsecure
}

// GetApiKey returns the ApiKey field.
func (a *AppLogDestinationSpecDataDog) GetApiKey() string {
	if a == nil {
		return ""
	}
	return a.ApiKey
}

// GetEndpoint returns the Endpoint field.
func (a *AppLogDestinationSpecDataDog) GetEndpoint() string {
	if a == nil {
		return ""
	}
	return a.Endpoint
}

// GetKey returns the Key field.
func (a *AppLogDestinationSpecHeader) GetKey() string {
	if a == nil {
		return ""
	}
	return a.Key
}

// GetValue returns the Value field.
func (a *AppLogDestinationSpecHeader) GetValue() string {
	if a == nil {
		return ""
	}
	return a.Value
}

// GetToken returns the Token field.
func (a *AppLogDestinationSpecLogtail) GetToken() string {
	if a == nil {
		return ""
	}
	return a.Token
}

// GetEndpoint returns the Endpoint field.
func (a *AppLogDestinationSpecPapertrail) GetEndpoint() string {
	if a == nil {
		return ""
	}
	return a.Endpoint
}

// GetAppID returns the AppID field.
func (a *AppProposeRequest) GetAppID() string {
	if a == nil {
		return ""
	}
	return a.AppID
}

// GetSpec returns the Spec field.
func (a *AppProposeRequest) GetSpec() *AppSpec {
	if a == nil {
		return nil
	}
	return a.Spec
}

// GetAppCost returns the AppCost field.
func (a *AppProposeResponse) GetAppCost() float32 {
	if a == nil {
		return 0
	}
	return a.AppCost
}

// GetAppIsStarter returns the AppIsStarter field.
func (a *AppProposeResponse) GetAppIsStarter() bool {
	if a == nil {
		return false
	}
	return a.AppIsStarter
}

// GetAppIsStatic returns the AppIsStatic field.
func (a *AppProposeResponse) GetAppIsStatic() bool {
	if a == nil {
		return false
	}
	return a.AppIsStatic
}

// GetAppNameAvailable returns the AppNameAvailable field.
func (a *AppProposeResponse) GetAppNameAvailable() bool {
	if a == nil {
		return false
	}
	return a.AppNameAvailable
}

// GetAppNameSuggestion returns the AppNameSuggestion field.
func (a *AppProposeResponse) GetAppNameSuggestion() string {
	if a == nil {
		return ""
	}
	return a.AppNameSuggestion
}

// GetAppTierDowngradeCost returns the AppTierDowngradeCost field.
func (a *AppProposeResponse) GetAppTierDowngradeCost() float32 {
	if a == nil {
		return 0
	}
	return a.AppTierDowngradeCost
}

// GetAppTierUpgradeCost returns the AppTierUpgradeCost field.
func (a *AppProposeResponse) GetAppTierUpgradeCost() float32 {
	if a == nil {
		return 0
	}
	return a.AppTierUpgradeCost
}

// GetExistingStarterApps returns the ExistingStarterApps field.
func (a *AppProposeResponse) GetExistingStarterApps() string {
	if a == nil {
		return ""
	}
	return a.ExistingStarterApps
}

// GetExistingStaticApps returns the ExistingStaticApps field.
func (a *AppProposeResponse) GetExistingStaticApps() string {
	if a == nil {
		return ""
	}
	return a.ExistingStaticApps
}

// GetMaxFreeStarterApps returns the MaxFreeStarterApps field.
func (a *AppProposeResponse) GetMaxFreeStarterApps() string {
	if a == nil {
		return ""
	}
	return a.MaxFreeStarterApps
}

// GetMaxFreeStaticApps returns the MaxFreeStaticApps field.
func (a *AppProposeResponse) GetMaxFreeStaticApps() string {
	if a == nil {
		return ""
	}
	return a.MaxFreeStaticApps
}

// GetSpec returns the Spec field.
func (a *AppProposeResponse) GetSpec() *AppSpec {
	if a == nil {
		return nil
	}
	return a.Spec
}

// GetContinent returns the Continent field.
func (a *AppRegion) GetContinent() string {
	if a == nil {
		return ""
	}
	return a.Continent
}

// GetDataCenters returns the DataCenters field.
func (a *AppRegion) GetDataCenters() []string {
	if a == nil {
		return nil
	}
	return a.DataCenters
}

// GetDefault returns the Default field.
func (a *AppRegion) GetDefault() bool {
	if a == nil {
		return false
	}
	return a.Default
}

// GetDisabled returns the Disabled field.
func (a *AppRegion) GetDisabled() bool {
	if a == nil {
		return false
	}
	return a.Disabled
}

// GetFlag returns the Flag field.
func (a *AppRegion) GetFlag() string {
	if a == nil {
		return ""
	}
	return a.Flag
}

// GetLabel returns the Label field.
func (a *AppRegion) GetLabel() string {
	if a == nil {
		return ""
	}
	return a.Label
}

// GetReason returns the Reason field.
func (a *AppRegion) GetReason() string {
	if a == nil {
		return ""
	}
	return a.Reason
}

// GetSlug returns the Slug field.
func (a *AppRegion) GetSlug() string {
	if a == nil {
		return ""
	}
	return a.Slug
}

// GetPath returns the Path field.
func (a *AppRouteSpec) GetPath() string {
	if a == nil {
		return ""
	}
	return a.Path
}

// GetPreservePathPrefix returns the PreservePathPrefix field.
func (a *AppRouteSpec) GetPreservePathPrefix() bool {
	if a == nil {
		return false
	}
	return a.PreservePathPrefix
}

// GetAlerts returns the Alerts field.
func (a *AppServiceSpec) GetAlerts() []*AppAlertSpec {
	if a == nil {
		return nil
	}
	return a.Alerts
}

// GetBuildCommand returns the BuildCommand field.
func (a *AppServiceSpec) GetBuildCommand() string {
	if a == nil {
		return ""
	}
	return a.BuildCommand
}

// GetCORS returns the CORS field.
func (a *AppServiceSpec) GetCORS() *AppCORSPolicy {
	if a == nil {
		return nil
	}
	return a.CORS
}

// GetDockerfilePath returns the DockerfilePath field.
func (a *AppServiceSpec) GetDockerfilePath() string {
	if a == nil {
		return ""
	}
	return a.DockerfilePath
}

// GetEnvironmentSlug returns the EnvironmentSlug field.
func (a *AppServiceSpec) GetEnvironmentSlug() string {
	if a == nil {
		return ""
	}
	return a.EnvironmentSlug
}

// GetEnvs returns the Envs field.
func (a *AppServiceSpec) GetEnvs() []*AppVariableDefinition {
	if a == nil {
		return nil
	}
	return a.Envs
}

// GetGit returns the Git field.
func (a *AppServiceSpec) GetGit() *GitSourceSpec {
	if a == nil {
		return nil
	}
	return a.Git
}

// GetGitHub returns the GitHub field.
func (a *AppServiceSpec) GetGitHub() *GitHubSourceSpec {
	if a == nil {
		return nil
	}
	return a.GitHub
}

// GetGitLab returns the GitLab field.
func (a *AppServiceSpec) GetGitLab() *GitLabSourceSpec {
	if a == nil {
		return nil
	}
	return a.GitLab
}

// GetHealthCheck returns the HealthCheck field.
func (a *AppServiceSpec) GetHealthCheck() *AppServiceSpecHealthCheck {
	if a == nil {
		return nil
	}
	return a.HealthCheck
}

// GetHTTPPort returns the HTTPPort field.
func (a *AppServiceSpec) GetHTTPPort() int64 {
	if a == nil {
		return 0
	}
	return a.HTTPPort
}

// GetImage returns the Image field.
func (a *AppServiceSpec) GetImage() *ImageSourceSpec {
	if a == nil {
		return nil
	}
	return a.Image
}

// GetInstanceCount returns the InstanceCount field.
func (a *AppServiceSpec) GetInstanceCount() int64 {
	if a == nil {
		return 0
	}
	return a.InstanceCount
}

// GetInstanceSizeSlug returns the InstanceSizeSlug field.
func (a *AppServiceSpec) GetInstanceSizeSlug() string {
	if a == nil {
		return ""
	}
	return a.InstanceSizeSlug
}

// GetInternalPorts returns the InternalPorts field.
func (a *AppServiceSpec) GetInternalPorts() []int64 {
	if a == nil {
		return nil
	}
	return a.InternalPorts
}

// GetLogDestinations returns the LogDestinations field.
func (a *AppServiceSpec) GetLogDestinations() []*AppLogDestinationSpec {
	if a == nil {
		return nil
	}
	return a.LogDestinations
}

// GetName returns the Name field.
func (a *AppServiceSpec) GetName() string {
	if a == nil {
		return ""
	}
	return a.Name
}

// GetRoutes returns the Routes field.
func (a *AppServiceSpec) GetRoutes() []*AppRouteSpec {
	if a == nil {
		return nil
	}
	return a.Routes
}

// GetRunCommand returns the RunCommand field.
func (a *AppServiceSpec) GetRunCommand() string {
	if a == nil {
		return ""
	}
	return a.RunCommand
}

// GetSourceDir returns the SourceDir field.
func (a *AppServiceSpec) GetSourceDir() string {
	if a == nil {
		return ""
	}
	return a.SourceDir
}

// GetFailureThreshold returns the FailureThreshold field.
func (a *AppServiceSpecHealthCheck) GetFailureThreshold() int32 {
	if a == nil {
		return 0
	}
	return a.FailureThreshold
}

// GetHTTPPath returns the HTTPPath field.
func (a *AppServiceSpecHealthCheck) GetHTTPPath() string {
	if a == nil {
		return ""
	}
	return a.HTTPPath
}

// GetInitialDelaySeconds returns the InitialDelaySeconds field.
func (a *AppServiceSpecHealthCheck) GetInitialDelaySeconds() int32 {
	if a == nil {
		return 0
	}
	return a.InitialDelaySeconds
}

// GetPath returns the Path field.
func (a *AppServiceSpecHealthCheck) GetPath() string {
	if a == nil {
		return ""
	}
	return a.Path
}

// GetPeriodSeconds returns the PeriodSeconds field.
func (a *AppServiceSpecHealthCheck) GetPeriodSeconds() int32 {
	if a == nil {
		return 0
	}
	return a.PeriodSeconds
}

// GetPort returns the Port field.
func (a *AppServiceSpecHealthCheck) GetPort() int64 {
	if a == nil {
		return 0
	}
	return a.Port
}

// GetSuccessThreshold returns the SuccessThreshold field.
func (a *AppServiceSpecHealthCheck) GetSuccessThreshold() int32 {
	if a == nil {
		return 0
	}
	return a.SuccessThreshold
}

// GetTimeoutSeconds returns the TimeoutSeconds field.
func (a *AppServiceSpecHealthCheck) GetTimeoutSeconds() int32 {
	if a == nil {
		return 0
	}
	return a.TimeoutSeconds
}

// GetAlerts returns the Alerts field.
func (a *AppSpec) GetAlerts() []*AppAlertSpec {
	if a == nil {
		return nil
	}
	return a.Alerts
}

// GetDatabases returns the Databases field.
func (a *AppSpec) GetDatabases() []*AppDatabaseSpec {
	if a == nil {
		return nil
	}
	return a.Databases
}

// GetDomains returns the Domains field.
func (a *AppSpec) GetDomains() []*AppDomainSpec {
	if a == nil {
		return nil
	}
	return a.Domains
}

// GetEnvs returns the Envs field.
func (a *AppSpec) GetEnvs() []*AppVariableDefinition {
	if a == nil {
		return nil
	}
	return a.Envs
}

// GetFeatures returns the Features field.
func (a *AppSpec) GetFeatures() []string {
	if a == nil {
		return nil
	}
	return a.Features
}

// GetFunctions returns the Functions field.
func (a *AppSpec) GetFunctions() []*AppFunctionsSpec {
	if a == nil {
		return nil
	}
	return a.Functions
}

// GetIngress returns the Ingress field.
func (a *AppSpec) GetIngress() *AppIngressSpec {
	if a == nil {
		return nil
	}
	return a.Ingress
}

// GetJobs returns the Jobs field.
func (a *AppSpec) GetJobs() []*AppJobSpec {
	if a == nil {
		return nil
	}
	return a.Jobs
}

// GetName returns the Name field.
func (a *AppSpec) GetName() string {
	if a == nil {
		return ""
	}
	return a.Name
}

// GetRegion returns the Region field.
func (a *AppSpec) GetRegion() string {
	if a == nil {
		return ""
	}
	return a.Region
}

// GetServices returns the Services field.
func (a *AppSpec) GetServices() []*AppServiceSpec {
	if a == nil {
		return nil
	}
	return a.Services
}

// GetStaticSites returns the StaticSites field.
func (a *AppSpec) GetStaticSites() []*AppStaticSiteSpec {
	if a == nil {
		return nil
	}
	return a.StaticSites
}

// GetWorkers returns the Workers field.
func (a *AppSpec) GetWorkers() []*AppWorkerSpec {
	if a == nil {
		return nil
	}
	return a.Workers
}

// GetBuildCommand returns the BuildCommand field.
func (a *AppStaticSiteSpec) GetBuildCommand() string {
	if a == nil {
		return ""
	}
	return a.BuildCommand
}

// GetCatchallDocument returns the CatchallDocument field.
func (a *AppStaticSiteSpec) GetCatchallDocument() string {
	if a == nil {
		return ""
	}
	return a.CatchallDocument
}

// GetCORS returns the CORS field.
func (a *AppStaticSiteSpec) GetCORS() *AppCORSPolicy {
	if a == nil {
		return nil
	}
	return a.CORS
}

// GetDockerfilePath returns the DockerfilePath field.
func (a *AppStaticSiteSpec) GetDockerfilePath() string {
	if a == nil {
		return ""
	}
	return a.DockerfilePath
}

// GetEnvironmentSlug returns the EnvironmentSlug field.
func (a *AppStaticSiteSpec) GetEnvironmentSlug() string {
	if a == nil {
		return ""
	}
	return a.EnvironmentSlug
}

// GetEnvs returns the Envs field.
func (a *AppStaticSiteSpec) GetEnvs() []*AppVariableDefinition {
	if a == nil {
		return nil
	}
	return a.Envs
}

// GetErrorDocument returns the ErrorDocument field.
func (a *AppStaticSiteSpec) GetErrorDocument() string {
	if a == nil {
		return ""
	}
	return a.ErrorDocument
}

// GetGit returns the Git field.
func (a *AppStaticSiteSpec) GetGit() *GitSourceSpec {
	if a == nil {
		return nil
	}
	return a.Git
}

// GetGitHub returns the GitHub field.
func (a *AppStaticSiteSpec) GetGitHub() *GitHubSourceSpec {
	if a == nil {
		return nil
	}
	return a.GitHub
}

// GetGitLab returns the GitLab field.
func (a *AppStaticSiteSpec) GetGitLab() *GitLabSourceSpec {
	if a == nil {
		return nil
	}
	return a.GitLab
}

// GetIndexDocument returns the IndexDocument field.
func (a *AppStaticSiteSpec) GetIndexDocument() string {
	if a == nil {
		return ""
	}
	return a.IndexDocument
}

// GetName returns the Name field.
func (a *AppStaticSiteSpec) GetName() string {
	if a == nil {
		return ""
	}
	return a.Name
}

// GetOutputDir returns the OutputDir field.
func (a *AppStaticSiteSpec) GetOutputDir() string {
	if a == nil {
		return ""
	}
	return a.OutputDir
}

// GetRoutes returns the Routes field.
func (a *AppStaticSiteSpec) GetRoutes() []*AppRouteSpec {
	if a == nil {
		return nil
	}
	return a.Routes
}

// GetSourceDir returns the SourceDir field.
func (a *AppStaticSiteSpec) GetSourceDir() string {
	if a == nil {
		return ""
	}
	return a.SourceDir
}

// GetExact returns the Exact field.
func (a *AppStringMatch) GetExact() string {
	if a == nil {
		return ""
	}
	return a.Exact
}

// GetPrefix returns the Prefix field.
func (a *AppStringMatch) GetPrefix() string {
	if a == nil {
		return ""
	}
	return a.Prefix
}

// GetRegex returns the Regex field.
func (a *AppStringMatch) GetRegex() string {
	if a == nil {
		return ""
	}
	return a.Regex
}

// GetBuildSeconds returns the BuildSeconds field.
func (a *AppTier) GetBuildSeconds() string {
	if a == nil {
		return ""
	}
	return a.BuildSeconds
}

// GetEgressBandwidthBytes returns the EgressBandwidthBytes field.
func (a *AppTier) GetEgressBandwidthBytes() string {
	if a == nil {
		return ""
	}
	return a.EgressBandwidthBytes
}

// GetName returns the Name field.
func (a *AppTier) GetName() string {
	if a == nil {
		return ""
	}
	return a.Name
}

// GetSlug returns the Slug field.
func (a *AppTier) GetSlug() string {
	if a == nil {
		return ""
	}
	return a.Slug
}

// GetKey returns the Key field.
func (a *AppVariableDefinition) GetKey() string {
	if a == nil {
		return ""
	}
	return a.Key
}

// GetScope returns the Scope field.
func (a *AppVariableDefinition) GetScope() AppVariableScope {
	if a == nil {
		return ""
	}
	return a.Scope
}

// GetType returns the Type field.
func (a *AppVariableDefinition) GetType() AppVariableType {
	if a == nil {
		return ""
	}
	return a.Type
}

// GetValue returns the Value field.
func (a *AppVariableDefinition) GetValue() string {
	if a == nil {
		return ""
	}
	return a.Value
}

// GetAlerts returns the Alerts field.
func (a *AppWorkerSpec) GetAlerts() []*AppAlertSpec {
	if a == nil {
		return nil
	}
	return a.Alerts
}

// GetBuildCommand returns the BuildCommand field.
func (a *AppWorkerSpec) GetBuildCommand() string {
	if a == nil {
		return ""
	}
	return a.BuildCommand
}

// GetDockerfilePath returns the DockerfilePath field.
func (a *AppWorkerSpec) GetDockerfilePath() string {
	if a == nil {
		return ""
	}
	return a.DockerfilePath
}

// GetEnvironmentSlug returns the EnvironmentSlug field.
func (a *AppWorkerSpec) GetEnvironmentSlug() string {
	if a == nil {
		return ""
	}
	return a.EnvironmentSlug
}

// GetEnvs returns the Envs field.
func (a *AppWorkerSpec) GetEnvs() []*AppVariableDefinition {
	if a == nil {
		return nil
	}
	return a.Envs
}

// GetGit returns the Git field.
func (a *AppWorkerSpec) GetGit() *GitSourceSpec {
	if a == nil {
		return nil
	}
	return a.Git
}

// GetGitHub returns the GitHub field.
func (a *AppWorkerSpec) GetGitHub() *GitHubSourceSpec {
	if a == nil {
		return nil
	}
	return a.GitHub
}

// GetGitLab returns the GitLab field.
func (a *AppWorkerSpec) GetGitLab() *GitLabSourceSpec {
	if a == nil {
		return nil
	}
	return a.GitLab
}

// GetImage returns the Image field.
func (a *AppWorkerSpec) GetImage() *ImageSourceSpec {
	if a == nil {
		return nil
	}
	return a.Image
}

// GetInstanceCount returns the InstanceCount field.
func (a *AppWorkerSpec) GetInstanceCount() int64 {
	if a == nil {
		return 0
	}
	return a.InstanceCount
}

// GetInstanceSizeSlug returns the InstanceSizeSlug field.
func (a *AppWorkerSpec) GetInstanceSizeSlug() string {
	if a == nil {
		return ""
	}
	return a.InstanceSizeSlug
}

// GetLogDestinations returns the LogDestinations field.
func (a *AppWorkerSpec) GetLogDestinations() []*AppLogDestinationSpec {
	if a == nil {
		return nil
	}
	return a.LogDestinations
}

// GetName returns the Name field.
func (a *AppWorkerSpec) GetName() string {
	if a == nil {
		return ""
	}
	return a.Name
}

// GetRunCommand returns the RunCommand field.
func (a *AppWorkerSpec) GetRunCommand() string {
	if a == nil {
		return ""
	}
	return a.RunCommand
}

// GetSourceDir returns the SourceDir field.
func (a *AppWorkerSpec) GetSourceDir() string {
	if a == nil {
		return ""
	}
	return a.SourceDir
}

// GetDescription returns the Description field.
func (b *Buildpack) GetDescription() []string {
	if b == nil {
		return nil
	}
	return b.Description
}

// GetDocsLink returns the DocsLink field.
func (b *Buildpack) GetDocsLink() string {
	if b == nil {
		return ""
	}
	return b.DocsLink
}

// GetID returns the ID field.
func (b *Buildpack) GetID() string {
	if b == nil {
		return ""
	}
	return b.ID
}

// GetLatest returns the Latest field.
func (b *Buildpack) GetLatest() bool {
	if b == nil {
		return false
	}
	return b.Latest
}

// GetMajorVersion returns the MajorVersion field.
func (b *Buildpack) GetMajorVersion() int32 {
	if b == nil {
		return 0
	}
	return b.MajorVersion
}

// GetName returns the Name field.
func (b *Buildpack) GetName() string {
	if b == nil {
		return ""
	}
	return b.Name
}

// GetVersion returns the Version field.
func (b *Buildpack) GetVersion() string {
	if b == nil {
		return ""
	}
	return b.Version
}

// GetCause returns the Cause field.
func (d *Deployment) GetCause() string {
	if d == nil {
		return ""
	}
	return d.Cause
}

// GetCauseDetails returns the CauseDetails field.
func (d *Deployment) GetCauseDetails() *DeploymentCauseDetails {
	if d == nil {
		return nil
	}
	return d.CauseDetails
}

// GetClonedFrom returns the ClonedFrom field.
func (d *Deployment) GetClonedFrom() string {
	if d == nil {
		return ""
	}
	return d.ClonedFrom
}

// GetCreatedAt returns the CreatedAt field.
func (d *Deployment) GetCreatedAt() time.Time {
	if d == nil {
		return time.Time{}
	}
	return d.CreatedAt
}

// GetFunctions returns the Functions field.
func (d *Deployment) GetFunctions() []*DeploymentFunctions {
	if d == nil {
		return nil
	}
	return d.Functions
}

// GetID returns the ID field.
func (d *Deployment) GetID() string {
	if d == nil {
		return ""
	}
	return d.ID
}

// GetJobs returns the Jobs field.
func (d *Deployment) GetJobs() []*DeploymentJob {
	if d == nil {
		return nil
	}
	return d.Jobs
}

// GetLoadBalancerID returns the LoadBalancerID field.
func (d *Deployment) GetLoadBalancerID() string {
	if d == nil {
		return ""
	}
	return d.LoadBalancerID
}

// GetPhase returns the Phase field.
func (d *Deployment) GetPhase() DeploymentPhase {
	if d == nil {
		return ""
	}
	return d.Phase
}

// GetPhaseLastUpdatedAt returns the PhaseLastUpdatedAt field.
func (d *Deployment) GetPhaseLastUpdatedAt() time.Time {
	if d == nil {
		return time.Time{}
	}
	return d.PhaseLastUpdatedAt
}

// GetPreviousDeploymentID returns the PreviousDeploymentID field.
func (d *Deployment) GetPreviousDeploymentID() string {
	if d == nil {
		return ""
	}
	return d.PreviousDeploymentID
}

// GetProgress returns the Progress field.
func (d *Deployment) GetProgress() *DeploymentProgress {
	if d == nil {
		return nil
	}
	return d.Progress
}

// GetServices returns the Services field.
func (d *Deployment) GetServices() []*DeploymentService {
	if d == nil {
		return nil
	}
	return d.Services
}

// GetSpec returns the Spec field.
func (d *Deployment) GetSpec() *AppSpec {
	if d == nil {
		return nil
	}
	return d.Spec
}

// GetStaticSites returns the StaticSites field.
func (d *Deployment) GetStaticSites() []*DeploymentStaticSite {
	if d == nil {
		return nil
	}
	return d.StaticSites
}

// GetTierSlug returns the TierSlug field.
func (d *Deployment) GetTierSlug() string {
	if d == nil {
		return ""
	}
	return d.TierSlug
}

// GetTiming returns the Timing field.
func (d *Deployment) GetTiming() *DeploymentTiming {
	if d == nil {
		return nil
	}
	return d.Timing
}

// GetUpdatedAt returns the UpdatedAt field.
func (d *Deployment) GetUpdatedAt() time.Time {
	if d == nil {
		return time.Time{}
	}
	return d.UpdatedAt
}

// GetWorkers returns the Workers field.
func (d *Deployment) GetWorkers() []*DeploymentWorker {
	if d == nil {
		return nil
	}
	return d.Workers
}

// GetDigitalOceanUserAction returns the DigitalOceanUserAction field.
func (d *DeploymentCauseDetails) GetDigitalOceanUserAction() *DeploymentCauseDetailsDigitalOceanUserAction {
	if d == nil {
		return nil
	}
	return d.DigitalOceanUserAction
}

// GetDOCRPush returns the DOCRPush field.
func (d *DeploymentCauseDetails) GetDOCRPush() *DeploymentCauseDetailsDOCRPush {
	if d == nil {
		return nil
	}
	return d.DOCRPush
}

// GetGitPush returns the GitPush field.
func (d *DeploymentCauseDetails) GetGitPush() *DeploymentCauseDetailsGitPush {
	if d == nil {
		return nil
	}
	return d.GitPush
}

// GetInternal returns the Internal field.
func (d *DeploymentCauseDetails) GetInternal() bool {
	if d == nil {
		return false
	}
	return d.Internal
}

// GetType returns the Type field.
func (d *DeploymentCauseDetails) GetType() DeploymentCauseDetailsType {
	if d == nil {
		return ""
	}
	return d.Type
}

// GetEmail returns the Email field.
func (d *DeploymentCauseDetailsDigitalOceanUser) GetEmail() string {
	if d == nil {
		return ""
	}
	return d.Email
}

// GetFullName returns the FullName field.
func (d *DeploymentCauseDetailsDigitalOceanUser) GetFullName() string {
	if d == nil {
		return ""
	}
	return d.FullName
}

// GetUUID returns the UUID field.
func (d *DeploymentCauseDetailsDigitalOceanUser) GetUUID() string {
	if d == nil {
		return ""
	}
	return d.UUID
}

// GetName returns the Name field.
func (d *DeploymentCauseDetailsDigitalOceanUserAction) GetName() DeploymentCauseDetailsDigitalOceanUserActionName {
	if d == nil {
		return ""
	}
	return d.Name
}

// GetUser returns the User field.
func (d *DeploymentCauseDetailsDigitalOceanUserAction) GetUser() *DeploymentCauseDetailsDigitalOceanUser {
	if d == nil {
		return nil
	}
	return d.User
}

// GetImageDigest returns the ImageDigest field.
func (d *DeploymentCauseDetailsDOCRPush) GetImageDigest() string {
	if d == nil {
		return ""
	}
	return d.ImageDigest
}

// GetRegistry returns the Registry field.
func (d *DeploymentCauseDetailsDOCRPush) GetRegistry() string {
	if d == nil {
		return ""
	}
	return d.Registry
}

// GetRepository returns the Repository field.
func (d *DeploymentCauseDetailsDOCRPush) GetRepository() string {
	if d == nil {
		return ""
	}
	return d.Repository
}

// GetTag returns the Tag field.
func (d *DeploymentCauseDetailsDOCRPush) GetTag() string {
	if d == nil {
		return ""
	}
	return d.Tag
}

// GetCommitAuthor returns the CommitAuthor field.
func (d *DeploymentCauseDetailsGitPush) GetCommitAuthor() string {
	if d == nil {
		return ""
	}
	return d.CommitAuthor
}

// GetCommitMessage returns the CommitMessage field.
func (d *DeploymentCauseDetailsGitPush) GetCommitMessage() string {
	if d == nil {
		return ""
	}
	return d.CommitMessage
}

// GetCommitSHA returns the CommitSHA field.
func (d *DeploymentCauseDetailsGitPush) GetCommitSHA() string {
	if d == nil {
		return ""
	}
	return d.CommitSHA
}

// GetGitHub returns the GitHub field.
func (d *DeploymentCauseDetailsGitPush) GetGitHub() *GitHubSourceSpec {
	if d == nil {
		return nil
	}
	return d.GitHub
}

// GetGitLab returns the GitLab field.
func (d *DeploymentCauseDetailsGitPush) GetGitLab() *GitLabSourceSpec {
	if d == nil {
		return nil
	}
	return d.GitLab
}

// GetUsername returns the Username field.
func (d *DeploymentCauseDetailsGitPush) GetUsername() string {
	if d == nil {
		return ""
	}
	return d.Username
}

// GetName returns the Name field.
func (d *DeploymentFunctions) GetName() string {
	if d == nil {
		return ""
	}
	return d.Name
}

// GetNamespace returns the Namespace field.
func (d *DeploymentFunctions) GetNamespace() string {
	if d == nil {
		return ""
	}
	return d.Namespace
}

// GetSourceCommitHash returns the SourceCommitHash field.
func (d *DeploymentFunctions) GetSourceCommitHash() string {
	if d == nil {
		return ""
	}
	return d.SourceCommitHash
}

// GetBuildpacks returns the Buildpacks field.
func (d *DeploymentJob) GetBuildpacks() []*Buildpack {
	if d == nil {
		return nil
	}
	return d.Buildpacks
}

// GetName returns the Name field.
func (d *DeploymentJob) GetName() string {
	if d == nil {
		return ""
	}
	return d.Name
}

// GetSourceCommitHash returns the SourceCommitHash field.
func (d *DeploymentJob) GetSourceCommitHash() string {
	if d == nil {
		return ""
	}
	return d.SourceCommitHash
}

// GetErrorSteps returns the ErrorSteps field.
func (d *DeploymentProgress) GetErrorSteps() int32 {
	if d == nil {
		return 0
	}
	return d.ErrorSteps
}

// GetPendingSteps returns the PendingSteps field.
func (d *DeploymentProgress) GetPendingSteps() int32 {
	if d == nil {
		return 0
	}
	return d.PendingSteps
}

// GetRunningSteps returns the RunningSteps field.
func (d *DeploymentProgress) GetRunningSteps() int32 {
	if d == nil {
		return 0
	}
	return d.RunningSteps
}

// GetSteps returns the Steps field.
func (d *DeploymentProgress) GetSteps() []*DeploymentProgressStep {
	if d == nil {
		return nil
	}
	return d.Steps
}

// GetSuccessSteps returns the SuccessSteps field.
func (d *DeploymentProgress) GetSuccessSteps() int32 {
	if d == nil {
		return 0
	}
	return d.SuccessSteps
}

// GetSummarySteps returns the SummarySteps field.
func (d *DeploymentProgress) GetSummarySteps() []*DeploymentProgressStep {
	if d == nil {
		return nil
	}
	return d.SummarySteps
}

// GetTotalSteps returns the TotalSteps field.
func (d *DeploymentProgress) GetTotalSteps() int32 {
	if d == nil {
		return 0
	}
	return d.TotalSteps
}

// GetComponentName returns the ComponentName field.
func (d *DeploymentProgressStep) GetComponentName() string {
	if d == nil {
		return ""
	}
	return d.ComponentName
}

// GetEndedAt returns the EndedAt field.
func (d *DeploymentProgressStep) GetEndedAt() time.Time {
	if d == nil {
		return time.Time{}
	}
	return d.EndedAt
}

// GetMessageBase returns the MessageBase field.
func (d *DeploymentProgressStep) GetMessageBase() string {
	if d == nil {
		return ""
	}
	return d.MessageBase
}

// GetName returns the Name field.
func (d *DeploymentProgressStep) GetName() string {
	if d == nil {
		return ""
	}
	return d.Name
}

// GetReason returns the Reason field.
func (d *DeploymentProgressStep) GetReason() *DeploymentProgressStepReason {
	if d == nil {
		return nil
	}
	return d.Reason
}

// GetStartedAt returns the StartedAt field.
func (d *DeploymentProgressStep) GetStartedAt() time.Time {
	if d == nil {
		return time.Time{}
	}
	return d.StartedAt
}

// GetStatus returns the Status field.
func (d *DeploymentProgressStep) GetStatus() DeploymentProgressStepStatus {
	if d == nil {
		return ""
	}
	return d.Status
}

// GetSteps returns the Steps field.
func (d *DeploymentProgressStep) GetSteps() []*DeploymentProgressStep {
	if d == nil {
		return nil
	}
	return d.Steps
}

// GetCode returns the Code field.
func (d *DeploymentProgressStepReason) GetCode() string {
	if d == nil {
		return ""
	}
	return d.Code
}

// GetMessage returns the Message field.
func (d *DeploymentProgressStepReason) GetMessage() string {
	if d == nil {
		return ""
	}
	return d.Message
}

// GetBuildpacks returns the Buildpacks field.
func (d *DeploymentService) GetBuildpacks() []*Buildpack {
	if d == nil {
		return nil
	}
	return d.Buildpacks
}

// GetName returns the Name field.
func (d *DeploymentService) GetName() string {
	if d == nil {
		return ""
	}
	return d.Name
}

// GetSourceCommitHash returns the SourceCommitHash field.
func (d *DeploymentService) GetSourceCommitHash() string {
	if d == nil {
		return ""
	}
	return d.SourceCommitHash
}

// GetBuildpacks returns the Buildpacks field.
func (d *DeploymentStaticSite) GetBuildpacks() []*Buildpack {
	if d == nil {
		return nil
	}
	return d.Buildpacks
}

// GetName returns the Name field.
func (d *DeploymentStaticSite) GetName() string {
	if d == nil {
		return ""
	}
	return d.Name
}

// GetSourceCommitHash returns the SourceCommitHash field.
func (d *DeploymentStaticSite) GetSourceCommitHash() string {
	if d == nil {
		return ""
	}
	return d.SourceCommitHash
}

// GetBuildBillable returns the BuildBillable field.
func (d *DeploymentTiming) GetBuildBillable() string {
	if d == nil {
		return ""
	}
	return d.BuildBillable
}

// GetBuildTotal returns the BuildTotal field.
func (d *DeploymentTiming) GetBuildTotal() string {
	if d == nil {
		return ""
	}
	return d.BuildTotal
}

// GetComponents returns the Components field.
func (d *DeploymentTiming) GetComponents() []*DeploymentTimingComponent {
	if d == nil {
		return nil
	}
	return d.Components
}

// GetDatabaseProvision returns the DatabaseProvision field.
func (d *DeploymentTiming) GetDatabaseProvision() string {
	if d == nil {
		return ""
	}
	return d.DatabaseProvision
}

// GetDeploying returns the Deploying field.
func (d *DeploymentTiming) GetDeploying() string {
	if d == nil {
		return ""
	}
	return d.Deploying
}

// GetPending returns the Pending field.
func (d *DeploymentTiming) GetPending() string {
	if d == nil {
		return ""
	}
	return d.Pending
}

// GetBuildBillable returns the BuildBillable field.
func (d *DeploymentTimingComponent) GetBuildBillable() string {
	if d == nil {
		return ""
	}
	return d.BuildBillable
}

// GetName returns the Name field.
func (d *DeploymentTimingComponent) GetName() string {
	if d == nil {
		return ""
	}
	return d.Name
}

// GetBuildpacks returns the Buildpacks field.
func (d *DeploymentWorker) GetBuildpacks() []*Buildpack {
	if d == nil {
		return nil
	}
	return d.Buildpacks
}

// GetName returns the Name field.
func (d *DeploymentWorker) GetName() string {
	if d == nil {
		return ""
	}
	return d.Name
}

// GetSourceCommitHash returns the SourceCommitHash field.
func (d *DeploymentWorker) GetSourceCommitHash() string {
	if d == nil {
		return ""
	}
	return d.SourceCommitHash
}

// GetSpec returns the Spec field.
func (d *DeployTemplate) GetSpec() *AppSpec {
	if d == nil {
		return nil
	}
	return d.Spec
}

// GetCommitSHA returns the CommitSHA field.
func (d *DetectRequest) GetCommitSHA() string {
	if d == nil {
		return ""
	}
	return d.CommitSHA
}

// GetGit returns the Git field.
func (d *DetectRequest) GetGit() *GitSourceSpec {
	if d == nil {
		return nil
	}
	return d.Git
}

// GetGitHub returns the GitHub field.
func (d *DetectRequest) GetGitHub() *GitHubSourceSpec {
	if d == nil {
		return nil
	}
	return d.GitHub
}

// GetGitLab returns the GitLab field.
func (d *DetectRequest) GetGitLab() *GitLabSourceSpec {
	if d == nil {
		return nil
	}
	return d.GitLab
}

// GetSourceDir returns the SourceDir field.
func (d *DetectRequest) GetSourceDir() string {
	if d == nil {
		return ""
	}
	return d.SourceDir
}

// GetComponents returns the Components field.
func (d *DetectResponse) GetComponents() []*DetectResponseComponent {
	if d == nil {
		return nil
	}
	return d.Components
}

// GetTemplate returns the Template field.
func (d *DetectResponse) GetTemplate() *DeployTemplate {
	if d == nil {
		return nil
	}
	return d.Template
}

// GetTemplateError returns the TemplateError field.
func (d *DetectResponse) GetTemplateError() string {
	if d == nil {
		return ""
	}
	return d.TemplateError
}

// GetTemplateFound returns the TemplateFound field.
func (d *DetectResponse) GetTemplateFound() bool {
	if d == nil {
		return false
	}
	return d.TemplateFound
}

// GetTemplateValid returns the TemplateValid field.
func (d *DetectResponse) GetTemplateValid() bool {
	if d == nil {
		return false
	}
	return d.TemplateValid
}

// GetBuildCommand returns the BuildCommand field.
func (d *DetectResponseComponent) GetBuildCommand() string {
	if d == nil {
		return ""
	}
	return d.BuildCommand
}

// GetBuildpacks returns the Buildpacks field.
func (d *DetectResponseComponent) GetBuildpacks() []*Buildpack {
	if d == nil {
		return nil
	}
	return d.Buildpacks
}

// GetDockerfiles returns the Dockerfiles field.
func (d *DetectResponseComponent) GetDockerfiles() []string {
	if d == nil {
		return nil
	}
	return d.Dockerfiles
}

// GetEnvironmentSlug returns the EnvironmentSlug field.
func (d *DetectResponseComponent) GetEnvironmentSlug() string {
	if d == nil {
		return ""
	}
	return d.EnvironmentSlug
}

// GetEnvVars returns the EnvVars field.
func (d *DetectResponseComponent) GetEnvVars() []*AppVariableDefinition {
	if d == nil {
		return nil
	}
	return d.EnvVars
}

// GetHTTPPorts returns the HTTPPorts field.
func (d *DetectResponseComponent) GetHTTPPorts() []int64 {
	if d == nil {
		return nil
	}
	return d.HTTPPorts
}

// GetRunCommand returns the RunCommand field.
func (d *DetectResponseComponent) GetRunCommand() string {
	if d == nil {
		return ""
	}
	return d.RunCommand
}

// GetServerlessPackages returns the ServerlessPackages field.
func (d *DetectResponseComponent) GetServerlessPackages() []*DetectResponseServerlessPackage {
	if d == nil {
		return nil
	}
	return d.ServerlessPackages
}

// GetSourceDir returns the SourceDir field.
func (d *DetectResponseComponent) GetSourceDir() string {
	if d == nil {
		return ""
	}
	return d.SourceDir
}

// GetStrategy returns the Strategy field.
func (d *DetectResponseComponent) GetStrategy() DetectResponseType {
	if d == nil {
		return ""
	}
	return d.Strategy
}

// GetTypes returns the Types field.
func (d *DetectResponseComponent) GetTypes() []string {
	if d == nil {
		return nil
	}
	return d.Types
}

// GetLimits returns the Limits field.
func (d *DetectResponseServerlessFunction) GetLimits() *DetectResponseServerlessFunctionLimits {
	if d == nil {
		return nil
	}
	return d.Limits
}

// GetName returns the Name field.
func (d *DetectResponseServerlessFunction) GetName() string {
	if d == nil {
		return ""
	}
	return d.Name
}

// GetPackage returns the Package field.
func (d *DetectResponseServerlessFunction) GetPackage() string {
	if d == nil {
		return ""
	}
	return d.Package
}

// GetRuntime returns the Runtime field.
func (d *DetectResponseServerlessFunction) GetRuntime() string {
	if d == nil {
		return ""
	}
	return d.Runtime
}

// GetLogs returns the Logs field.
func (d *DetectResponseServerlessFunctionLimits) GetLogs() string {
	if d == nil {
		return ""
	}
	return d.Logs
}

// GetMemory returns the Memory field.
func (d *DetectResponseServerlessFunctionLimits) GetMemory() string {
	if d == nil {
		return ""
	}
	return d.Memory
}

// GetTimeout returns the Timeout field.
func (d *DetectResponseServerlessFunctionLimits) GetTimeout() string {
	if d == nil {
		return ""
	}
	return d.Timeout
}

// GetFunctions returns the Functions field.
func (d *DetectResponseServerlessPackage) GetFunctions() []*DetectResponseServerlessFunction {
	if d == nil {
		return nil
	}
	return d.Functions
}

// GetName returns the Name field.
func (d *DetectResponseServerlessPackage) GetName() string {
	if d == nil {
		return ""
	}
	return d.Name
}

// GetBranch returns the Branch field.
func (g *GitHubSourceSpec) GetBranch() string {
	if g == nil {
		return ""
	}
	return g.Branch
}

// GetDeployOnPush returns the DeployOnPush field.
func (g *GitHubSourceSpec) GetDeployOnPush() bool {
	if g == nil {
		return false
	}
	return g.DeployOnPush
}

// GetRepo returns the Repo field.
func (g *GitHubSourceSpec) GetRepo() string {
	if g == nil {
		return ""
	}
	return g.Repo
}

// GetBranch returns the Branch field.
func (g *GitLabSourceSpec) GetBranch() string {
	if g == nil {
		return ""
	}
	return g.Branch
}

// GetDeployOnPush returns the DeployOnPush field.
func (g *GitLabSourceSpec) GetDeployOnPush() bool {
	if g == nil {
		return false
	}
	return g.DeployOnPush
}

// GetRepo returns the Repo field.
func (g *GitLabSourceSpec) GetRepo() string {
	if g == nil {
		return ""
	}
	return g.Repo
}

// GetBranch returns the Branch field.
func (g *GitSourceSpec) GetBranch() string {
	if g == nil {
		return ""
	}
	return g.Branch
}

// GetRepoCloneURL returns the RepoCloneURL field.
func (g *GitSourceSpec) GetRepoCloneURL() string {
	if g == nil {
		return ""
	}
	return g.RepoCloneURL
}

// GetDeployOnPush returns the DeployOnPush field.
func (i *ImageSourceSpec) GetDeployOnPush() *ImageSourceSpecDeployOnPush {
	if i == nil {
		return nil
	}
	return i.DeployOnPush
}

// GetDigest returns the Digest field.
func (i *ImageSourceSpec) GetDigest() string {
	if i == nil {
		return ""
	}
	return i.Digest
}

// GetRegistry returns the Registry field.
func (i *ImageSourceSpec) GetRegistry() string {
	if i == nil {
		return ""
	}
	return i.Registry
}

// GetRegistryType returns the RegistryType field.
func (i *ImageSourceSpec) GetRegistryType() ImageSourceSpecRegistryType {
	if i == nil {
		return ""
	}
	return i.RegistryType
}

// GetRepository returns the Repository field.
func (i *ImageSourceSpec) GetRepository() string {
	if i == nil {
		return ""
	}
	return i.Repository
}

// GetTag returns the Tag field.
func (i *ImageSourceSpec) GetTag() string {
	if i == nil {
		return ""
	}
	return i.Tag
}

// GetEnabled returns the Enabled field.
func (i *ImageSourceSpecDeployOnPush) GetEnabled() bool {
	if i == nil {
		return false
	}
	return i.Enabled
}

// GetBuildpacks returns the Buildpacks field.
func (l *ListBuildpacksResponse) GetBuildpacks() []*Buildpack {
	if l == nil {
		return nil
	}
	return l.Buildpacks
}

// GetAffectedComponents returns the AffectedComponents field.
func (u *UpgradeBuildpackResponse) GetAffectedComponents() []string {
	if u == nil {
		return nil
	}
	return u.AffectedComponents
}

// GetDeployment returns the Deployment field.
func (u *UpgradeBuildpackResponse) GetDeployment() *Deployment {
	if u == nil {
		return nil
	}
	return u.Deployment
}
