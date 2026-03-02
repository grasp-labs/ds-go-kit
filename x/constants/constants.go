package constants

import "regexp"

// ================== Validation patterns ==================
const (
	RegPatternNotAlphaNumUnderscore string = `[^a-z0-9_]`
	RegPatternStringSearch          string = `(?i)^[a-z0-9., _-]*$`
	RegPatternAWSS3BucketName       string = `^([a-z0-9]+(-[a-z0-9]+)*){3,63}$`
	RegPatternAWSRegion             string = `^[a-z]{2}-[a-z]+-[0-9]+$`
	RegPatternAWSS3Prefix           string = `^([a-zA-Z0-9!_.*'()-]+/)*$`
	RegPatternDPName                string = `^[a-zA-Z0-9æøåÆØÅ\s]+$`
	RegPatternDPServiceName         string = `^[A-Za-z\-]+$`
	RegPatternProductID             string = `^[A-Za-z0-9](?:[A-Za-z0-9-]*[A-Za-z0-9])?\$` +
		`[A-Za-z0-9](?:[A-Za-z0-9-]*[A-Za-z0-9])?\$` +
		`[A-Za-z0-9](?:[A-Za-z0-9-]*[A-Za-z0-9])?\$` +
		`(?:[vV])?(?:0|[1-9]\d*)(?:\.(?:0|[1-9]\d*)){2,3}$`
	RegPatternTopicID string = `^` +
		`([a-z0-9][a-z0-9-]*)\.` +
		`([a-z0-9][a-z0-9-]*)\.` +
		`([a-z0-9][a-z0-9-]*)` +
		`(?:\.([a-z0-9][a-z0-9-]*))?` +
		`\.` +
		`([a-z0-9][a-z0-9-]*)` +
		`\.v` +
		`([1-9]\d*)` +
		`$`
)

// ================== Compiled regular expressions ==================
var RegExpNotAlphaNumUnderscore = regexp.MustCompile(RegPatternNotAlphaNumUnderscore)
var RegExpStringSearch = regexp.MustCompile(RegPatternStringSearch)
var RegExpAWSBucketName = regexp.MustCompile(RegPatternAWSS3BucketName)
var RegExpAWSRegion = regexp.MustCompile(RegPatternAWSRegion)
var RegExpAWSS3Prefix = regexp.MustCompile(RegPatternAWSS3Prefix)
var RegExpDPName = regexp.MustCompile(RegPatternDPName)
var RegExpDPServiceName = regexp.MustCompile(RegPatternDPServiceName)
var RegExpDPProductID = regexp.MustCompile(RegPatternProductID)
var RegExpDPTopicID = regexp.MustCompile(RegPatternTopicID)

// ================== Default values ==================
const DefaultLanguage = "norwegian"
