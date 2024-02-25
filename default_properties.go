// Code generated, DO NOT EDIT.

package hubspot

type ContactDefaultProperties struct {
	// Contact's company size. Required for the Facebook Ads Integration. Automatically synced from the Lead Ads tool.
	CompanySize string `json:"company_size,omitempty"`

	// Contact's date of birth. Required for the Facebook Ads Integration. Automatically synced from the Lead Ads tool.
	DateOfBirth string `json:"date_of_birth,omitempty"`

	// Count of days elapsed between creation and being closed as a customer. Set automatically.
	DaysToClose Int `json:"days_to_close,omitempty"`

	// Contact's degree. Required for the Facebook Ads Integration. Automatically synced from the Lead Ads tool.
	Degree string `json:"degree,omitempty"`

	// Contact's field of study. Required for the Facebook Ads Integration. Automatically synced from the Lead Ads tool.
	FieldOfStudy string `json:"field_of_study,omitempty"`

	// Date  this  contact  first  submitted  a  form.
	FirstConversionDate DateTime `json:"first_conversion_date,omitempty"`

	// First form this contact submitted.
	FirstConversionEventName string `json:"first_conversion_event_name,omitempty"`

	// Date first deal was created for contact. Set automatically.
	FirstDealCreatedDate DateTime `json:"first_deal_created_date,omitempty"`

	// Contact's gender. Required for the Facebook Ads Integration. Automatically synced from the Lead Ads tool.
	Gender string `json:"gender,omitempty"`

	// Contact's graduation date. Required for the Facebook Ads Integration. Automatically synced from the Lead Ads tool.
	GraduationDate string `json:"graduation_date,omitempty"`

	// A set of additional email addresses for a contact
	HsAdditionalEmails Enumeration `json:"hs_additional_emails,omitempty"`

	// The business units this record is assigned to.
	HsAllAssignedBusinessUnitIds Enumeration `json:"hs_all_assigned_business_unit_ids,omitempty"`

	// A set of all vids, canonical or otherwise, for a contact
	HsAllContactVids Enumeration `json:"hs_all_contact_vids,omitempty"`

	// Campaign responsible for the first touch creation of this contact.
	HsAnalyticsFirstTouchConvertingCampaign string `json:"hs_analytics_first_touch_converting_campaign,omitempty"`

	// Campaign responsible for the last touch creation of this contact.
	HsAnalyticsLastTouchConvertingCampaign string `json:"hs_analytics_last_touch_converting_campaign,omitempty"`

	// The path in the FileManager CDN for this contact's avatar override image. Automatically set by HubSpot.
	HsAvatarFilemanagerKey string `json:"hs_avatar_filemanager_key,omitempty"`

	// Role the contact plays during the sales process. Contacts can have multiple roles, and can share roles with others.
	HsBuyingRole Enumeration `json:"hs_buying_role,omitempty"`

	// A set of all form submissions for a contact
	HsCalculatedFormSubmissions Enumeration `json:"hs_calculated_form_submissions,omitempty"`

	// Merged vids with timestamps of a contact
	HsCalculatedMergedVids Enumeration `json:"hs_calculated_merged_vids,omitempty"`

	// Mobile number in international format
	HsCalculatedMobileNumber string `json:"hs_calculated_mobile_number,omitempty"`

	// Phone number in international format
	HsCalculatedPhoneNumber string `json:"hs_calculated_phone_number,omitempty"`

	// Area Code of the calculated phone number
	HsCalculatedPhoneNumberAreaCode string `json:"hs_calculated_phone_number_area_code,omitempty"`

	// Country code of the calculated phone number
	HsCalculatedPhoneNumberCountryCode string `json:"hs_calculated_phone_number_country_code,omitempty"`

	// ISO2 Country code for the derived phone number
	HsCalculatedPhoneNumberRegionCode string `json:"hs_calculated_phone_number_region_code,omitempty"`

	// Whether contact has clicked on a LinkedIn Ad
	HsClickedLinkedinAd Enumeration `json:"hs_clicked_linkedin_ad,omitempty"`

	// Email used to send private content information to members
	HsContentMembershipEmail string `json:"hs_content_membership_email,omitempty"`

	// Email Confirmation status of user of Content Membership
	HsContentMembershipEmailConfirmed Bool `json:"hs_content_membership_email_confirmed,omitempty"`

	// The time when the contact was first enrolled in the registration follow up email flow
	HsContentMembershipFollowUpEnqueuedAt DateTime `json:"hs_content_membership_follow_up_enqueued_at,omitempty"`

	// Notes relating to the contact's content membership.
	HsContentMembershipNotes string `json:"hs_content_membership_notes,omitempty"`

	// Datetime at which this user was set up for Content Membership
	HsContentMembershipRegisteredAt DateTime `json:"hs_content_membership_registered_at,omitempty"`

	// Domain to which the registration invitation email for Content Membership was sent to
	HsContentMembershipRegistrationDomainSentTo string `json:"hs_content_membership_registration_domain_sent_to,omitempty"`

	// Datetime at which this user was sent a registration invitation email for Content Membership
	HsContentMembershipRegistrationEmailSentAt DateTime `json:"hs_content_membership_registration_email_sent_at,omitempty"`

	// Status of the contact's content membership.
	HsContentMembershipStatus Enumeration `json:"hs_content_membership_status,omitempty"`

	// A Conversations visitor's email address
	HsConversationsVisitorEmail string `json:"hs_conversations_visitor_email,omitempty"`

	// if contact is assigned and unworked, set to 1. if contact is assigned and worked, set to 0.
	HsCountIsUnworked Int `json:"hs_count_is_unworked,omitempty"`

	// if contact is assigned and worked, set to 1. if contact is assigned and unworked, set to 0.
	HsCountIsWorked Int `json:"hs_count_is_worked,omitempty"`

	// Flag indicating this contact was created by the Conversations API
	HsCreatedByConversations Bool `json:"hs_created_by_conversations,omitempty"`

	// The user that created this object. This value is automatically set by HubSpot and may not be modified.
	HsCreatedByUserId Int `json:"hs_created_by_user_id,omitempty"`

	// The date and time at which this object was created. This value is automatically set by HubSpot and may not be modified.
	HsCreatedate DateTime `json:"hs_createdate,omitempty"`

	// This property captures ads consents from forms and is used by consentmanager to create / update associated data privacy consent objects
	HsDataPrivacyAdsConsent Bool `json:"hs_data_privacy_ads_consent,omitempty"`

	// The date and time when the contact entered the 'Customer' stage, 'Lifecycle Stage Pipeline' pipeline
	HsDateEnteredCustomer DateTime `json:"hs_date_entered_customer,omitempty"`

	// The date and time when the contact entered the 'Evangelist' stage, 'Lifecycle Stage Pipeline' pipeline
	HsDateEnteredEvangelist DateTime `json:"hs_date_entered_evangelist,omitempty"`

	// The date and time when the contact entered the 'Lead' stage, 'Lifecycle Stage Pipeline' pipeline
	HsDateEnteredLead DateTime `json:"hs_date_entered_lead,omitempty"`

	// The date and time when the contact entered the 'Marketing Qualified Lead' stage, 'Lifecycle Stage Pipeline' pipeline
	HsDateEnteredMarketingqualifiedlead DateTime `json:"hs_date_entered_marketingqualifiedlead,omitempty"`

	// The date and time when the contact entered the 'Opportunity' stage, 'Lifecycle Stage Pipeline' pipeline
	HsDateEnteredOpportunity DateTime `json:"hs_date_entered_opportunity,omitempty"`

	// The date and time when the contact entered the 'Other' stage, 'Lifecycle Stage Pipeline' pipeline
	HsDateEnteredOther DateTime `json:"hs_date_entered_other,omitempty"`

	// The date and time when the contact entered the 'Sales Qualified Lead' stage, 'Lifecycle Stage Pipeline' pipeline
	HsDateEnteredSalesqualifiedlead DateTime `json:"hs_date_entered_salesqualifiedlead,omitempty"`

	// The date and time when the contact entered the 'Subscriber' stage, 'Lifecycle Stage Pipeline' pipeline
	HsDateEnteredSubscriber DateTime `json:"hs_date_entered_subscriber,omitempty"`

	// The date and time when the contact exited the 'Customer' stage, 'Lifecycle Stage Pipeline' pipeline
	HsDateExitedCustomer DateTime `json:"hs_date_exited_customer,omitempty"`

	// The date and time when the contact exited the 'Evangelist' stage, 'Lifecycle Stage Pipeline' pipeline
	HsDateExitedEvangelist DateTime `json:"hs_date_exited_evangelist,omitempty"`

	// The date and time when the contact exited the 'Lead' stage, 'Lifecycle Stage Pipeline' pipeline
	HsDateExitedLead DateTime `json:"hs_date_exited_lead,omitempty"`

	// The date and time when the contact exited the 'Marketing Qualified Lead' stage, 'Lifecycle Stage Pipeline' pipeline
	HsDateExitedMarketingqualifiedlead DateTime `json:"hs_date_exited_marketingqualifiedlead,omitempty"`

	// The date and time when the contact exited the 'Opportunity' stage, 'Lifecycle Stage Pipeline' pipeline
	HsDateExitedOpportunity DateTime `json:"hs_date_exited_opportunity,omitempty"`

	// The date and time when the contact exited the 'Other' stage, 'Lifecycle Stage Pipeline' pipeline
	HsDateExitedOther DateTime `json:"hs_date_exited_other,omitempty"`

	// The date and time when the contact exited the 'Sales Qualified Lead' stage, 'Lifecycle Stage Pipeline' pipeline
	HsDateExitedSalesqualifiedlead DateTime `json:"hs_date_exited_salesqualifiedlead,omitempty"`

	// The date and time when the contact exited the 'Subscriber' stage, 'Lifecycle Stage Pipeline' pipeline
	HsDateExitedSubscriber DateTime `json:"hs_date_exited_subscriber,omitempty"`

	// The last time a shared document (presentation) was accessed by this contact
	HsDocumentLastRevisited DateTime `json:"hs_document_last_revisited,omitempty"`

	// The email address associated with this contact is invalid.
	HsEmailBadAddress Bool `json:"hs_email_bad_address,omitempty"`

	// The reason why the email address has been quarantined.
	HsEmailCustomerQuarantinedReason Enumeration `json:"hs_email_customer_quarantined_reason,omitempty"`

	// A contact's email address domain
	HsEmailDomain string `json:"hs_email_domain,omitempty"`

	// The issue that caused a contact to hard bounce from your emails. If this is an error or a temporary issue, you can unbounce this contact from the contact record.
	HsEmailHardBounceReason string `json:"hs_email_hard_bounce_reason,omitempty"`

	// The issue that caused a contact to hard bounce from your emails. If this is an error or a temporary issue, you can unbounce this contact from the contact record.
	HsEmailHardBounceReasonEnum Enumeration `json:"hs_email_hard_bounce_reason_enum,omitempty"`

	// Indicates that the current email address has been quarantined for anti-abuse reasons and any marketing email sends to it will be blocked. This is automatically set by HubSpot.
	HsEmailQuarantined Bool `json:"hs_email_quarantined,omitempty"`

	// The automated reason why the email address has been quarantined.
	HsEmailQuarantinedReason Enumeration `json:"hs_email_quarantined_reason,omitempty"`

	// When this recipient has reached the limit of email sends per time period, this property indicates the next available time to send. This is automatically set by HubSpot.
	HsEmailRecipientFatigueRecoveryTime DateTime `json:"hs_email_recipient_fatigue_recovery_time,omitempty"`

	// The number of marketing emails that have been sent to the current email address since the last engagement (open or link click). This is automatically set by HubSpot.
	HsEmailSendsSinceLastEngagement Int `json:"hs_email_sends_since_last_engagement,omitempty"`

	// The status of a contact's eligibility to receive marketing email. This is automatically set by HubSpot.
	HsEmailconfirmationstatus Enumeration `json:"hs_emailconfirmationstatus,omitempty"`

	// Whether contact has clicked a Facebook ad
	HsFacebookAdClicked Bool `json:"hs_facebook_ad_clicked,omitempty"`

	//
	HsFacebookClickId string `json:"hs_facebook_click_id,omitempty"`

	// A contact's facebook id
	HsFacebookid string `json:"hs_facebookid,omitempty"`

	// Last NPS survey comment that this contact gave
	HsFeedbackLastNpsFollowUp string `json:"hs_feedback_last_nps_follow_up,omitempty"`

	// Last NPS survey rating that this contact gave
	HsFeedbackLastNpsRating Enumeration `json:"hs_feedback_last_nps_rating,omitempty"`

	// The time that this contact last submitted a NPS survey response. This is automatically set by HubSpot.
	HsFeedbackLastSurveyDate DateTime `json:"hs_feedback_last_survey_date,omitempty"`

	// Whether or not a contact should be shown an NPS web survey. This is automatically set by HubSpot.
	HsFeedbackShowNpsWebSurvey Bool `json:"hs_feedback_show_nps_web_survey,omitempty"`

	// The object id of the current contact owner's first engagement with the contact.
	HsFirstEngagementObjectId Int `json:"hs_first_engagement_object_id,omitempty"`

	// The date of the first outreach (call, email, meeting or other communication) from a sales rep to the contact.
	HsFirstOutreachDate DateTime `json:"hs_first_outreach_date,omitempty"`

	// The create date of the first subscription by the contact.
	HsFirstSubscriptionCreateDate DateTime `json:"hs_first_subscription_create_date,omitempty"`

	//
	HsGoogleClickId string `json:"hs_google_click_id,omitempty"`

	// A contact's googleplus id
	HsGoogleplusid string `json:"hs_googleplusid,omitempty"`

	// The rollup property value is 1 when the contact has an active Subscription or 0 otherwise.
	HsHasActiveSubscription Int `json:"hs_has_active_subscription,omitempty"`

	// The timezone reported by a contact's IP address. This is automatically set by HubSpot and can be used for segmentation and reporting.
	HsIpTimezone string `json:"hs_ip_timezone,omitempty"`

	// Is a contact, has not been deleted and is not a visitor
	HsIsContact Bool `json:"hs_is_contact,omitempty"`

	// Contact  has  not  been  assigned  or  has  not  been  engaged  after  last  owner  assignment/re-assignment
	HsIsUnworked Bool `json:"hs_is_unworked,omitempty"`

	// The date of the last sales activity with the contact. This property is set automatically by HubSpot. Note: This field is only updated for contacts with an owner.
	HsLastSalesActivityDate DateTime `json:"hs_last_sales_activity_date,omitempty"`

	// The last time a contact engaged with your site or a form, document, meetings link, or tracked email. This doesn't include marketing emails or emails to multiple contacts.
	HsLastSalesActivityTimestamp DateTime `json:"hs_last_sales_activity_timestamp,omitempty"`

	// The type of the last engagement a contact performed. This doesn't include marketing emails or emails to multiple contacts.
	HsLastSalesActivityType Enumeration `json:"hs_last_sales_activity_type,omitempty"`

	// Most recent timestamp of any property update for this contact. This includes HubSpot internal properties, which can be visible or hidden. This property is updated automatically.
	HsLastmodifieddate DateTime `json:"hs_lastmodifieddate,omitempty"`

	// The most recent time at which an associated lead currently in a disqualified stage was moved to that stage
	HsLatestDisqualifiedLeadDate DateTime `json:"hs_latest_disqualified_lead_date,omitempty"`

	// The most recent time an associated open lead was moved to a NEW or IN_PROGRESS state
	HsLatestOpenLeadDate DateTime `json:"hs_latest_open_lead_date,omitempty"`

	// The most recent time at which an associated lead currently in a qualified stage was moved to that stage
	HsLatestQualifiedLeadDate DateTime `json:"hs_latest_qualified_lead_date,omitempty"`

	// The last sequence ended date.
	HsLatestSequenceEndedDate DateTime `json:"hs_latest_sequence_ended_date,omitempty"`

	// The last sequence enrolled.
	HsLatestSequenceEnrolled Int `json:"hs_latest_sequence_enrolled,omitempty"`

	// The last sequence enrolled date.
	HsLatestSequenceEnrolledDate DateTime `json:"hs_latest_sequence_enrolled_date,omitempty"`

	// The last sequence finished date.
	HsLatestSequenceFinishedDate DateTime `json:"hs_latest_sequence_finished_date,omitempty"`

	// The last sequence unenrolled date.
	HsLatestSequenceUnenrolledDate DateTime `json:"hs_latest_sequence_unenrolled_date,omitempty"`

	// The time of the latest session for a contact
	HsLatestSourceTimestamp DateTime `json:"hs_latest_source_timestamp,omitempty"`

	// The create date of the latest subscription by the contact.
	HsLatestSubscriptionCreateDate DateTime `json:"hs_latest_subscription_create_date,omitempty"`

	// The contact's sales, prospecting or outreach status
	HsLeadStatus Enumeration `json:"hs_lead_status,omitempty"`

	// Legal basis for processing contact's data; 'Not applicable' will exempt the contact from GDPR protections
	HsLegalBasis Enumeration `json:"hs_legal_basis,omitempty"`

	//
	HsLinkedinAdClicked Enumeration `json:"hs_linkedin_ad_clicked,omitempty"`

	// A contact's linkedin id
	HsLinkedinid string `json:"hs_linkedinid,omitempty"`

	// The ID of the activity that set the contact as a marketing contact
	HsMarketableReasonId string `json:"hs_marketable_reason_id,omitempty"`

	// The type of the activity that set the contact as a marketing contact
	HsMarketableReasonType Enumeration `json:"hs_marketable_reason_type,omitempty"`

	// The marketing status of a contact
	HsMarketableStatus Enumeration `json:"hs_marketable_status,omitempty"`

	// Specifies if this contact will be set as non-marketing on renewal
	HsMarketableUntilRenewal Enumeration `json:"hs_marketable_until_renewal,omitempty"`

	// The list of Contact record IDs that have been merged into this Contact. This value is automatically set by HubSpot and may not be modified.
	HsMergedObjectIds Enumeration `json:"hs_merged_object_ids,omitempty"`

	// The unique ID for this record. This value is automatically set by HubSpot and may not be modified.
	HsObjectId Int64 `json:"hs_object_id,omitempty"`

	// Source (PropertySource) that created this object record
	HsObjectSource string `json:"hs_object_source,omitempty"`

	// First level of detail on how this record was created
	HsObjectSourceDetail1 string `json:"hs_object_source_detail_1,omitempty"`

	// Second level of detail on how this record was created
	HsObjectSourceDetail2 string `json:"hs_object_source_detail_2,omitempty"`

	// Third level of detail on how this record was created
	HsObjectSourceDetail3 string `json:"hs_object_source_detail_3,omitempty"`

	// The sourceId -- further detail -- of the source that created this object record
	HsObjectSourceId string `json:"hs_object_source_id,omitempty"`

	// How this record was created
	HsObjectSourceLabel Enumeration `json:"hs_object_source_label,omitempty"`

	// User ID of the user who initiated creation of this object record
	HsObjectSourceUserId Int `json:"hs_object_source_user_id,omitempty"`

	//
	HsPinnedEngagementId Int `json:"hs_pinned_engagement_id,omitempty"`

	// The pipeline with which this contact is currently associated
	HsPipeline Enumeration `json:"hs_pipeline,omitempty"`

	// The probability that a contact will become a customer within the next 90 days. This score is based on standard contact properties and behavior.
	HsPredictivecontactscoreV2 Int `json:"hs_predictivecontactscore_v2,omitempty"`

	// A ranking system of contacts evenly assigned into four tiers. Contacts in tier one are more likely to become customers than contacts in tier four.
	HsPredictivescoringtier Enumeration `json:"hs_predictivescoringtier,omitempty"`

	// Is the object read only
	HsReadOnly Bool `json:"hs_read_only,omitempty"`

	// The date the current contact owner first engaged with the contact.
	HsSaFirstEngagementDate DateTime `json:"hs_sa_first_engagement_date,omitempty"`

	// A description of the current contact owner's first engagement with the contact.
	HsSaFirstEngagementDescr Enumeration `json:"hs_sa_first_engagement_descr,omitempty"`

	// The object type of the current contact owner's first engagement with the contact.
	HsSaFirstEngagementObjectType Enumeration `json:"hs_sa_first_engagement_object_type,omitempty"`

	// The last time a tracked sales email was clicked by this user
	HsSalesEmailLastClicked DateTime `json:"hs_sales_email_last_clicked,omitempty"`

	// The last time a tracked sales email was opened by this contact. This property does not update for emails that were sent to more than one contact.
	HsSalesEmailLastOpened DateTime `json:"hs_sales_email_last_opened,omitempty"`

	// Mobile number with country code
	HsSearchableCalculatedInternationalMobileNumber string `json:"hs_searchable_calculated_international_mobile_number,omitempty"`

	// Phone number with country code
	HsSearchableCalculatedInternationalPhoneNumber string `json:"hs_searchable_calculated_international_phone_number,omitempty"`

	// Mobile number without country code
	HsSearchableCalculatedMobileNumber string `json:"hs_searchable_calculated_mobile_number,omitempty"`

	// Phone number without country code
	HsSearchableCalculatedPhoneNumber string `json:"hs_searchable_calculated_phone_number,omitempty"`

	// The number of sequences actively enrolled.
	HsSequencesActivelyEnrolledCount Int `json:"hs_sequences_actively_enrolled_count,omitempty"`

	// The number of sequences enrolled.
	HsSequencesEnrolledCount Int `json:"hs_sequences_enrolled_count,omitempty"`

	// A yes/no field that indicates whether the contact is currently in a Sequence.
	HsSequencesIsEnrolled Bool `json:"hs_sequences_is_enrolled,omitempty"`

	// The ID of the object from which the data was migrated. This is set automatically during portal data migration.
	HsSourceObjectId Int `json:"hs_source_object_id,omitempty"`

	// The ID of the portal from which the data was migrated. This is set automatically during portal data migration.
	HsSourcePortalId Int `json:"hs_source_portal_id,omitempty"`

	// testing purge
	HsTestpurge string `json:"hs_testpurge,omitempty"`

	// testing rollback
	HsTestrollback string `json:"hs_testrollback,omitempty"`

	//
	HsTimeBetweenContactCreationAndDealClose Int `json:"hs_time_between_contact_creation_and_deal_close,omitempty"`

	//
	HsTimeBetweenContactCreationAndDealCreation Int `json:"hs_time_between_contact_creation_and_deal_creation,omitempty"`

	// The total time in seconds spent by the contact in the 'Customer' stage, 'Lifecycle Stage Pipeline' pipeline
	HsTimeInCustomer Int `json:"hs_time_in_customer,omitempty"`

	// The total time in seconds spent by the contact in the 'Evangelist' stage, 'Lifecycle Stage Pipeline' pipeline
	HsTimeInEvangelist Int `json:"hs_time_in_evangelist,omitempty"`

	// The total time in seconds spent by the contact in the 'Lead' stage, 'Lifecycle Stage Pipeline' pipeline
	HsTimeInLead Int `json:"hs_time_in_lead,omitempty"`

	// The total time in seconds spent by the contact in the 'Marketing Qualified Lead' stage, 'Lifecycle Stage Pipeline' pipeline
	HsTimeInMarketingqualifiedlead Int `json:"hs_time_in_marketingqualifiedlead,omitempty"`

	// The total time in seconds spent by the contact in the 'Opportunity' stage, 'Lifecycle Stage Pipeline' pipeline
	HsTimeInOpportunity Int `json:"hs_time_in_opportunity,omitempty"`

	// The total time in seconds spent by the contact in the 'Other' stage, 'Lifecycle Stage Pipeline' pipeline
	HsTimeInOther Int `json:"hs_time_in_other,omitempty"`

	// The total time in seconds spent by the contact in the 'Sales Qualified Lead' stage, 'Lifecycle Stage Pipeline' pipeline
	HsTimeInSalesqualifiedlead Int `json:"hs_time_in_salesqualifiedlead,omitempty"`

	// The total time in seconds spent by the contact in the 'Subscriber' stage, 'Lifecycle Stage Pipeline' pipeline
	HsTimeInSubscriber Int `json:"hs_time_in_subscriber,omitempty"`

	// Time it took current owner to do first qualifying engagement.
	HsTimeToFirstEngagement Int `json:"hs_time_to_first_engagement,omitempty"`

	//  How long it takes for a contact to move from the HubSpot lead stage to the HubSpot customer stage.
	HsTimeToMoveFromLeadToCustomer Int `json:"hs_time_to_move_from_lead_to_customer,omitempty"`

	//  How long it takes for a contact to move from the HubSpot marketing qualified lead stage to the HubSpot customer stage.
	HsTimeToMoveFromMarketingqualifiedleadToCustomer Int `json:"hs_time_to_move_from_marketingqualifiedlead_to_customer,omitempty"`

	//  How long it takes for a contact to move from the HubSpot opportunity stage to the HubSpot customer stage.
	HsTimeToMoveFromOpportunityToCustomer Int `json:"hs_time_to_move_from_opportunity_to_customer,omitempty"`

	//  How long it takes for a contact to move from the HubSpot sales qualified lead stage to the HubSpot customer stage.
	HsTimeToMoveFromSalesqualifiedleadToCustomer Int `json:"hs_time_to_move_from_salesqualifiedlead_to_customer,omitempty"`

	//  How long it takes for a contact to move from the HubSpot subscriber stage to the HubSpot customer stage.
	HsTimeToMoveFromSubscriberToCustomer Int `json:"hs_time_to_move_from_subscriber_to_customer,omitempty"`

	// The contact’s time zone. This can be set automatically by HubSpot based on other contact properties. It can also be set manually for each contact.
	HsTimezone Enumeration `json:"hs_timezone,omitempty"`

	// A contact's twitter id
	HsTwitterid string `json:"hs_twitterid,omitempty"`

	// Unique property used for idempotent creates
	HsUniqueCreationKey string `json:"hs_unique_creation_key,omitempty"`

	// The user that last updated this object. This value is automatically set by HubSpot and may not be modified.
	HsUpdatedByUserId Int `json:"hs_updated_by_user_id,omitempty"`

	// The user IDs of all users that have clicked follow within the object to opt-in to getting follow notifications
	HsUserIdsOfAllNotificationFollowers Enumeration `json:"hs_user_ids_of_all_notification_followers,omitempty"`

	// The user IDs of all object owners that have clicked unfollow within the object to opt-out of getting follow notifications
	HsUserIdsOfAllNotificationUnfollowers Enumeration `json:"hs_user_ids_of_all_notification_unfollowers,omitempty"`

	// The user IDs of all owners of this object
	HsUserIdsOfAllOwners Enumeration `json:"hs_user_ids_of_all_owners,omitempty"`

	// The cumulative time in seconds spent by the contact in the 'Customer' stage, 'Lifecycle Stage Pipeline' pipeline
	HsV2CumulativeTimeInCustomer Int `json:"hs_v2_cumulative_time_in_customer,omitempty"`

	// The cumulative time in seconds spent by the contact in the 'Evangelist' stage, 'Lifecycle Stage Pipeline' pipeline
	HsV2CumulativeTimeInEvangelist Int `json:"hs_v2_cumulative_time_in_evangelist,omitempty"`

	// The cumulative time in seconds spent by the contact in the 'Lead' stage, 'Lifecycle Stage Pipeline' pipeline
	HsV2CumulativeTimeInLead Int `json:"hs_v2_cumulative_time_in_lead,omitempty"`

	// The cumulative time in seconds spent by the contact in the 'Marketing Qualified Lead' stage, 'Lifecycle Stage Pipeline' pipeline
	HsV2CumulativeTimeInMarketingqualifiedlead Int `json:"hs_v2_cumulative_time_in_marketingqualifiedlead,omitempty"`

	// The cumulative time in seconds spent by the contact in the 'Opportunity' stage, 'Lifecycle Stage Pipeline' pipeline
	HsV2CumulativeTimeInOpportunity Int `json:"hs_v2_cumulative_time_in_opportunity,omitempty"`

	// The cumulative time in seconds spent by the contact in the 'Other' stage, 'Lifecycle Stage Pipeline' pipeline
	HsV2CumulativeTimeInOther Int `json:"hs_v2_cumulative_time_in_other,omitempty"`

	// The cumulative time in seconds spent by the contact in the 'Sales Qualified Lead' stage, 'Lifecycle Stage Pipeline' pipeline
	HsV2CumulativeTimeInSalesqualifiedlead Int `json:"hs_v2_cumulative_time_in_salesqualifiedlead,omitempty"`

	// The cumulative time in seconds spent by the contact in the 'Subscriber' stage, 'Lifecycle Stage Pipeline' pipeline
	HsV2CumulativeTimeInSubscriber Int `json:"hs_v2_cumulative_time_in_subscriber,omitempty"`

	// The date and time when the contact entered the 'Customer' stage, 'Lifecycle Stage Pipeline' pipeline
	HsV2DateEnteredCustomer DateTime `json:"hs_v2_date_entered_customer,omitempty"`

	// The date and time when the contact entered the 'Evangelist' stage, 'Lifecycle Stage Pipeline' pipeline
	HsV2DateEnteredEvangelist DateTime `json:"hs_v2_date_entered_evangelist,omitempty"`

	// The date and time when the contact entered the 'Lead' stage, 'Lifecycle Stage Pipeline' pipeline
	HsV2DateEnteredLead DateTime `json:"hs_v2_date_entered_lead,omitempty"`

	// The date and time when the contact entered the 'Marketing Qualified Lead' stage, 'Lifecycle Stage Pipeline' pipeline
	HsV2DateEnteredMarketingqualifiedlead DateTime `json:"hs_v2_date_entered_marketingqualifiedlead,omitempty"`

	// The date and time when the contact entered the 'Opportunity' stage, 'Lifecycle Stage Pipeline' pipeline
	HsV2DateEnteredOpportunity DateTime `json:"hs_v2_date_entered_opportunity,omitempty"`

	// The date and time when the contact entered the 'Other' stage, 'Lifecycle Stage Pipeline' pipeline
	HsV2DateEnteredOther DateTime `json:"hs_v2_date_entered_other,omitempty"`

	// The date and time when the contact entered the 'Sales Qualified Lead' stage, 'Lifecycle Stage Pipeline' pipeline
	HsV2DateEnteredSalesqualifiedlead DateTime `json:"hs_v2_date_entered_salesqualifiedlead,omitempty"`

	// The date and time when the contact entered the 'Subscriber' stage, 'Lifecycle Stage Pipeline' pipeline
	HsV2DateEnteredSubscriber DateTime `json:"hs_v2_date_entered_subscriber,omitempty"`

	// The date and time when the contact exited the 'Customer' stage, 'Lifecycle Stage Pipeline' pipeline
	HsV2DateExitedCustomer DateTime `json:"hs_v2_date_exited_customer,omitempty"`

	// The date and time when the contact exited the 'Evangelist' stage, 'Lifecycle Stage Pipeline' pipeline
	HsV2DateExitedEvangelist DateTime `json:"hs_v2_date_exited_evangelist,omitempty"`

	// The date and time when the contact exited the 'Lead' stage, 'Lifecycle Stage Pipeline' pipeline
	HsV2DateExitedLead DateTime `json:"hs_v2_date_exited_lead,omitempty"`

	// The date and time when the contact exited the 'Marketing Qualified Lead' stage, 'Lifecycle Stage Pipeline' pipeline
	HsV2DateExitedMarketingqualifiedlead DateTime `json:"hs_v2_date_exited_marketingqualifiedlead,omitempty"`

	// The date and time when the contact exited the 'Opportunity' stage, 'Lifecycle Stage Pipeline' pipeline
	HsV2DateExitedOpportunity DateTime `json:"hs_v2_date_exited_opportunity,omitempty"`

	// The date and time when the contact exited the 'Other' stage, 'Lifecycle Stage Pipeline' pipeline
	HsV2DateExitedOther DateTime `json:"hs_v2_date_exited_other,omitempty"`

	// The date and time when the contact exited the 'Sales Qualified Lead' stage, 'Lifecycle Stage Pipeline' pipeline
	HsV2DateExitedSalesqualifiedlead DateTime `json:"hs_v2_date_exited_salesqualifiedlead,omitempty"`

	// The date and time when the contact exited the 'Subscriber' stage, 'Lifecycle Stage Pipeline' pipeline
	HsV2DateExitedSubscriber DateTime `json:"hs_v2_date_exited_subscriber,omitempty"`

	// The total time in seconds spent by the contact in the 'Customer' stage, 'Lifecycle Stage Pipeline' pipeline since it last entered this stage
	HsV2LatestTimeInCustomer Int `json:"hs_v2_latest_time_in_customer,omitempty"`

	// The total time in seconds spent by the contact in the 'Evangelist' stage, 'Lifecycle Stage Pipeline' pipeline since it last entered this stage
	HsV2LatestTimeInEvangelist Int `json:"hs_v2_latest_time_in_evangelist,omitempty"`

	// The total time in seconds spent by the contact in the 'Lead' stage, 'Lifecycle Stage Pipeline' pipeline since it last entered this stage
	HsV2LatestTimeInLead Int `json:"hs_v2_latest_time_in_lead,omitempty"`

	// The total time in seconds spent by the contact in the 'Marketing Qualified Lead' stage, 'Lifecycle Stage Pipeline' pipeline since it last entered this stage
	HsV2LatestTimeInMarketingqualifiedlead Int `json:"hs_v2_latest_time_in_marketingqualifiedlead,omitempty"`

	// The total time in seconds spent by the contact in the 'Opportunity' stage, 'Lifecycle Stage Pipeline' pipeline since it last entered this stage
	HsV2LatestTimeInOpportunity Int `json:"hs_v2_latest_time_in_opportunity,omitempty"`

	// The total time in seconds spent by the contact in the 'Other' stage, 'Lifecycle Stage Pipeline' pipeline since it last entered this stage
	HsV2LatestTimeInOther Int `json:"hs_v2_latest_time_in_other,omitempty"`

	// The total time in seconds spent by the contact in the 'Sales Qualified Lead' stage, 'Lifecycle Stage Pipeline' pipeline since it last entered this stage
	HsV2LatestTimeInSalesqualifiedlead Int `json:"hs_v2_latest_time_in_salesqualifiedlead,omitempty"`

	// The total time in seconds spent by the contact in the 'Subscriber' stage, 'Lifecycle Stage Pipeline' pipeline since it last entered this stage
	HsV2LatestTimeInSubscriber Int `json:"hs_v2_latest_time_in_subscriber,omitempty"`

	// Object is part of an import
	HsWasImported Bool `json:"hs_was_imported,omitempty"`

	// The phone number associated with the contact’s WhatsApp account.
	HsWhatsappPhoneNumber string `json:"hs_whatsapp_phone_number,omitempty"`

	// The most recent date a HubSpot Owner was assigned to a contact. This is set automatically by HubSpot and can be used for segmentation and reporting.
	HubspotOwnerAssigneddate DateTime `json:"hubspot_owner_assigneddate,omitempty"`

	// The city reported by a contact's IP address. This is automatically set by HubSpot and can be used for segmentation and reporting.
	IpCity string `json:"ip_city,omitempty"`

	// The country reported by a contact's IP address. This is automatically set by HubSpot and can be used for segmentation and reporting.
	IpCountry string `json:"ip_country,omitempty"`

	// The country code reported by a contact's IP address. This is automatically set by HubSpot and can be used for segmentation and reporting.
	IpCountryCode string `json:"ip_country_code,omitempty"`

	//
	IpLatlon string `json:"ip_latlon,omitempty"`

	// The state or region reported by a contact's IP address. This is automatically set by HubSpot and can be used for segmentation and reporting.
	IpState string `json:"ip_state,omitempty"`

	// The state code or region code reported by a contact's IP address. This is automatically set by HubSpot and can be used for segmentation and reporting.
	IpStateCode string `json:"ip_state_code,omitempty"`

	// The zipcode reported by a contact's IP address. This is automatically set by HubSpot and can be used for segmentation and reporting.
	IpZipcode string `json:"ip_zipcode,omitempty"`

	// Contact's job function. Required for the Facebook Ads Integration. Automatically synced from the Lead Ads tool.
	JobFunction string `json:"job_function,omitempty"`

	// The date any property on this contact was modified
	Lastmodifieddate DateTime `json:"lastmodifieddate,omitempty"`

	// Contact's marital status. Required for the Facebook Ads Integration. Automatically synced from the Lead Ads tool.
	MaritalStatus string `json:"marital_status,omitempty"`

	// Contact's military status. Required for the Facebook Ads Integration. Automatically synced from the Lead Ads tool.
	MilitaryStatus string `json:"military_status,omitempty"`

	// Count of deals associated with this contact. Set automatically by HubSpot.
	NumAssociatedDeals Int `json:"num_associated_deals,omitempty"`

	// The number of forms this contact has submitted
	NumConversionEvents Int `json:"num_conversion_events,omitempty"`

	// The number of different forms this contact has submitted
	NumUniqueConversionEvents Int `json:"num_unique_conversion_events,omitempty"`

	// The date this contact last submitted a form
	RecentConversionDate DateTime `json:"recent_conversion_date,omitempty"`

	// The last form this contact submitted
	RecentConversionEventName string `json:"recent_conversion_event_name,omitempty"`

	// Amount of last closed won deal associated with a contact. Set automatically.
	RecentDealAmount Int `json:"recent_deal_amount,omitempty"`

	// Date last deal associated with contact was closed-won. Set automatically.
	RecentDealCloseDate DateTime `json:"recent_deal_close_date,omitempty"`

	// Contact's relationship status. Required for the Facebook Ads Integration. Automatically synced from the Lead Ads tool.
	RelationshipStatus string `json:"relationship_status,omitempty"`

	// Contact's school. Required for the Facebook Ads Integration. Automatically synced from the Lead Ads tool.
	School string `json:"school,omitempty"`

	// Contact's seniority. Required for the Facebook Ads Integration. Automatically synced from the Lead Ads tool.
	Seniority string `json:"seniority,omitempty"`

	// Contact's start date. Required for the Facebook Ads Integration. Automatically synced from the Lead Ads tool.
	StartDate string `json:"start_date,omitempty"`

	// Sum of all closed-won deal revenue associated with the contact. Set automatically.
	TotalRevenue Int `json:"total_revenue,omitempty"`

	// Contact's work email. Required for the Facebook Ads Integration. Automatically synced from the Lead Ads tool.
	WorkEmail string `json:"work_email,omitempty"`

	// A contact's first name
	Firstname string `json:"firstname,omitempty"`

	// First page the contact visited on your website. Set automatically.
	HsAnalyticsFirstUrl string `json:"hs_analytics_first_url,omitempty"`

	// The number of marketing emails delivered for the current email address. This is automatically set by HubSpot.
	HsEmailDelivered Int `json:"hs_email_delivered,omitempty"`

	// Indicates that the current email address has opted out of this email type.
	HsEmailOptout299470044 Enumeration `json:"hs_email_optout_299470044,omitempty"`

	// Indicates that the current email address has opted out of this email type.
	HsEmailOptout299470045 Enumeration `json:"hs_email_optout_299470045,omitempty"`

	// Indicates that the current email address has opted out of this email type.
	HsEmailOptout299470046 Enumeration `json:"hs_email_optout_299470046,omitempty"`

	// The contact's Twitter handle.
	Twitterhandle string `json:"twitterhandle,omitempty"`

	// True when contact is enrolled in a workflow.
	Currentlyinworkflow Enumeration `json:"currentlyinworkflow,omitempty"`

	// The number of Twitter followers a contact has
	Followercount Int `json:"followercount,omitempty"`

	// Last page the contact visited on your website. Set automatically.
	HsAnalyticsLastUrl string `json:"hs_analytics_last_url,omitempty"`

	// The number of marketing emails opened for the current email address. This is automatically set by HubSpot.
	HsEmailOpen Int `json:"hs_email_open,omitempty"`

	// A contact's last name
	Lastname string `json:"lastname,omitempty"`

	// Total number of page views this contact has had on your website. Set automatically.
	HsAnalyticsNumPageViews Int `json:"hs_analytics_num_page_views,omitempty"`

	// The number of marketing emails which have had link clicks for the current email address. This is automatically set by HubSpot.
	HsEmailClick Int `json:"hs_email_click,omitempty"`

	// The title used to address a contact
	Salutation string `json:"salutation,omitempty"`

	// The contact's Twitter profile photo. This is set by HubSpot using the contact's email address.
	Twitterprofilephoto string `json:"twitterprofilephoto,omitempty"`

	// A contact's email address
	Email string `json:"email,omitempty"`

	// Number of times a contact has come to your website. Set automatically.
	HsAnalyticsNumVisits Int `json:"hs_analytics_num_visits,omitempty"`

	// The number of marketing emails that bounced for the current email address. This is automatically set by HubSpot.
	HsEmailBounce Int `json:"hs_email_bounce,omitempty"`

	// A contact's persona
	HsPersona Enumeration `json:"hs_persona,omitempty"`

	// The date of the most recent click on a published social message. This is set automatically by HubSpot for each contact.
	HsSocialLastEngagement DateTime `json:"hs_social_last_engagement,omitempty"`

	// Total number of events for this contact. Set automatically.
	HsAnalyticsNumEventCompletions Int `json:"hs_analytics_num_event_completions,omitempty"`

	// Indicates that the current email address has opted out of all email.
	HsEmailOptout Bool `json:"hs_email_optout,omitempty"`

	// The number of times a contact clicked on links you shared on Twitter through HubSpot. This is set automatically by HubSpot and can be used for segmentation.
	HsSocialTwitterClicks Int `json:"hs_social_twitter_clicks,omitempty"`

	// A contact's mobile phone number
	Mobilephone string `json:"mobilephone,omitempty"`

	// A contact's primary phone number
	Phone string `json:"phone,omitempty"`

	// A contact's primary fax number
	Fax string `json:"fax,omitempty"`

	// First time the contact has been seen. Set automatically.
	HsAnalyticsFirstTimestamp DateTime `json:"hs_analytics_first_timestamp,omitempty"`

	// The name of the last marketing email sent to the current email address. This is automatically set by HubSpot.
	HsEmailLastEmailName string `json:"hs_email_last_email_name,omitempty"`

	// The date of the most recent delivery for any marketing email to the current email address. This is automatically set by HubSpot.
	HsEmailLastSendDate DateTime `json:"hs_email_last_send_date,omitempty"`

	// The number clicks on links shared on Facebook
	HsSocialFacebookClicks Int `json:"hs_social_facebook_clicks,omitempty"`

	// Contact's street address, including apartment or unit number.
	Address string `json:"address,omitempty"`

	// The date of the last meeting that has been scheduled by a contact through the meetings tool. If multiple meetings have been scheduled, the date of the last chronological meeting in the timeline is shown.
	EngagementsLastMeetingBooked DateTime `json:"engagements_last_meeting_booked,omitempty"`

	// UTM parameter for marketing campaign (e.g. a specific email) responsible for recent meeting booking. Only populated when tracking parameters are included in meeting link.
	EngagementsLastMeetingBookedCampaign string `json:"engagements_last_meeting_booked_campaign,omitempty"`

	// UTM parameter for the channel (e.g. email)  responsible for most recent meeting booking. Only populated when tracking parameters are included in meeting link.
	EngagementsLastMeetingBookedMedium string `json:"engagements_last_meeting_booked_medium,omitempty"`

	// UTM parameter for the site (e.g. Twitter) responsible for most recent meeting booking. Only populated when tracking parameters are included in meeting link.
	EngagementsLastMeetingBookedSource string `json:"engagements_last_meeting_booked_source,omitempty"`

	// First time the contact visited your website. Set automatically.
	HsAnalyticsFirstVisitTimestamp DateTime `json:"hs_analytics_first_visit_timestamp,omitempty"`

	// The date of the most recent open for any marketing email to the current email address. This is automatically set by HubSpot.
	HsEmailLastOpenDate DateTime `json:"hs_email_last_open_date,omitempty"`

	// The date of the most recent meeting (past or upcoming) logged for, scheduled with, or booked by this contact.
	HsLatestMeetingActivity DateTime `json:"hs_latest_meeting_activity,omitempty"`

	// The last time a tracked sales email was replied to by this user. This is set automatically by HubSpot based on user actions in the contact record.
	HsSalesEmailLastReplied DateTime `json:"hs_sales_email_last_replied,omitempty"`

	// The number clicks on links shared on LinkedIn
	HsSocialLinkedinClicks Int `json:"hs_social_linkedin_clicks,omitempty"`

	// The owner of a contact. This can be any HubSpot user or Salesforce integration user, and can be set manually or via Workflows.
	HubspotOwnerId Enumeration `json:"hubspot_owner_id,omitempty"`

	// The last time a call, chat conversation, LinkedIn message, postal mail, meeting, sales email, SMS, or WhatsApp message was logged for a contact. This is set automatically by HubSpot based on user actions in the contact record.
	NotesLastContacted DateTime `json:"notes_last_contacted,omitempty"`

	// The last time a call, chat conversation, LinkedIn message, postal mail, meeting, note, sales email, SMS, or WhatsApp message was logged for a contact. This is set automatically by HubSpot based on user actions in the contact record.
	NotesLastUpdated DateTime `json:"notes_last_updated,omitempty"`

	// The date of the next upcoming activity for a contact. This is set automatically by HubSpot based on user actions in the contact record.
	NotesNextActivityDate DateTime `json:"notes_next_activity_date,omitempty"`

	// The number of times a call, chat conversation, LinkedIn message, postal mail, meeting, sales email, SMS, or WhatsApp message was logged for a contact record. This is set automatically by HubSpot based on user actions in the contact record.
	NumContactedNotes Int `json:"num_contacted_notes,omitempty"`

	// The number of times a call, chat conversation, LinkedIn message, postal mail, meeting, note, sales email, SMS, task, or WhatsApp message was logged for a contact record. This is set automatically by HubSpot based on user actions in the contact record.
	NumNotes Int `json:"num_notes,omitempty"`

	// A legacy property used to identify the email address of the owner of the contact. This property is no longer in use.
	Owneremail string `json:"owneremail,omitempty"`

	// A legacy property used to identify the name of the owner of the contact. This property is no longer in use.
	Ownername string `json:"ownername,omitempty"`

	// This field is meaningless on its own, and is solely used for triggering dynamic list updates when SurveyMonkey information is updated
	Surveymonkeyeventlastupdated Int `json:"surveymonkeyeventlastupdated,omitempty"`

	// This field is meaningless on its own, and is solely used for triggering dynamic list updates when webinar information is updated
	Webinareventlastupdated Int `json:"webinareventlastupdated,omitempty"`

	// A contact's city of residence
	City string `json:"city,omitempty"`

	// Timestamp for most recent webpage view on your website.
	HsAnalyticsLastTimestamp DateTime `json:"hs_analytics_last_timestamp,omitempty"`

	// The date of the most recent link click for any marketing email to the current email address. This is automatically set by HubSpot.
	HsEmailLastClickDate DateTime `json:"hs_email_last_click_date,omitempty"`

	// The number clicks on links shared on Google Plus
	HsSocialGooglePlusClicks Int `json:"hs_social_google_plus_clicks,omitempty"`

	// The team of the owner of a contact.
	HubspotTeamId Enumeration `json:"hubspot_team_id,omitempty"`

	// A contact's LinkedIn bio
	Linkedinbio string `json:"linkedinbio,omitempty"`

	// The contact's Twitter bio. This is set by HubSpot using the contact's email address.
	Twitterbio string `json:"twitterbio,omitempty"`

	// The value of all owner referencing properties for this object, both default and custom
	HsAllOwnerIds Enumeration `json:"hs_all_owner_ids,omitempty"`

	// Timestamp for start of the most recent session for this contact to your website.
	HsAnalyticsLastVisitTimestamp DateTime `json:"hs_analytics_last_visit_timestamp,omitempty"`

	// The date of the earliest delivery for any marketing email to the current email address. This is automatically set by HubSpot.
	HsEmailFirstSendDate DateTime `json:"hs_email_first_send_date,omitempty"`

	// The number of clicks on published social messages. This is set automatically by HubSpot for each contact.
	HsSocialNumBroadcastClicks Int `json:"hs_social_num_broadcast_clicks,omitempty"`

	// The contact's state of residence. This might be set via import, form, or integration.
	State string `json:"state,omitempty"`

	// The team ids corresponding to all owner referencing properties for this object, both default and custom
	HsAllTeamIds Enumeration `json:"hs_all_team_ids,omitempty"`

	// First known source the contact used to find your website. Set automatically, but may be updated manually.
	HsAnalyticsSource Enumeration `json:"hs_analytics_source,omitempty"`

	// The date of the earliest open for any marketing email to the current email address. This is automatically set by HubSpot.
	HsEmailFirstOpenDate DateTime `json:"hs_email_first_open_date,omitempty"`

	// The source of the latest session for a contact
	HsLatestSource Enumeration `json:"hs_latest_source,omitempty"`

	// The contact's zip code. This might be set via import, form, or integration.
	Zip string `json:"zip,omitempty"`

	// The contact's country/region of residence. This might be set via import, form, or integration.
	Country string `json:"country,omitempty"`

	// The team ids, including up the team hierarchy, corresponding to all owner referencing properties for this object, both default and custom
	HsAllAccessibleTeamIds Enumeration `json:"hs_all_accessible_team_ids,omitempty"`

	// Additional information about the source the contact used to find your website. Set automatically.
	HsAnalyticsSourceData1 string `json:"hs_analytics_source_data_1,omitempty"`

	// The date of the earliest link click for any marketing email to the current email address. This is automatically set by HubSpot.
	HsEmailFirstClickDate DateTime `json:"hs_email_first_click_date,omitempty"`

	// Additional information about the latest source for the last session the contact used to find your website. Set automatically.
	HsLatestSourceData1 string `json:"hs_latest_source_data_1,omitempty"`

	// How many LinkedIn connections they have
	Linkedinconnections Int `json:"linkedinconnections,omitempty"`

	// Additional information about the source the contact used to find your website. Set automatically.
	HsAnalyticsSourceData2 string `json:"hs_analytics_source_data_2,omitempty"`

	// Indicates the contact is globally ineligible for email.
	HsEmailIsIneligible Bool `json:"hs_email_is_ineligible,omitempty"`

	// Set your contact's preferred language for communications. This property can be changed from an import, form, or integration.
	HsLanguage Enumeration `json:"hs_language,omitempty"`

	// Additional information about the source for the last session the contact used to find your website. Set automatically.
	HsLatestSourceData2 string `json:"hs_latest_source_data_2,omitempty"`

	// A contact's Klout score, a measure of Internet influence
	Kloutscoregeneral Int `json:"kloutscoregeneral,omitempty"`

	// URL that referred the contact to your website. Set automatically.
	HsAnalyticsFirstReferrer string `json:"hs_analytics_first_referrer,omitempty"`

	// The date of the earliest reply for any marketing email to the current email address. This is automatically set by HubSpot.
	HsEmailFirstReplyDate DateTime `json:"hs_email_first_reply_date,omitempty"`

	// A contact's job title
	Jobtitle string `json:"jobtitle,omitempty"`

	// Social Media photo
	Photo string `json:"photo,omitempty"`

	// Last URL that referred contact to your website. Set automatically.
	HsAnalyticsLastReferrer string `json:"hs_analytics_last_referrer,omitempty"`

	// The date of the latest reply for any marketing email to the current email address. This is automatically set by HubSpot.
	HsEmailLastReplyDate DateTime `json:"hs_email_last_reply_date,omitempty"`

	// A default property to be used for any message or comments a contact may want to leave on a form.
	Message string `json:"message,omitempty"`

	// Date the  contact  became  a  customer. Set automatically when a deal or opportunity is marked as closed-won. It can also be set manually or programmatically.
	Closedate DateTime `json:"closedate,omitempty"`

	// Average number of pageviews per session for this contact. Set automatically.
	HsAnalyticsAveragePageViews Int `json:"hs_analytics_average_page_views,omitempty"`

	// The number of marketing emails replied to by the current email address. This is automatically set by HubSpot.
	HsEmailReplied Int `json:"hs_email_replied,omitempty"`

	// Set event revenue on a contact though the Enterprise Events feature. http://help.hubspot.com/articles/KCS_Article/Reports/How-do-I-create-Events-in-HubSpot
	HsAnalyticsRevenue Int `json:"hs_analytics_revenue,omitempty"`

	// The date when a contact's lifecycle stage changed to Lead. This is automatically set by HubSpot for each contact.
	HsLifecyclestageLeadDate DateTime `json:"hs_lifecyclestage_lead_date,omitempty"`

	// The date when a contact's lifecycle stage changed to MQL. This is automatically set by HubSpot for each contact.
	HsLifecyclestageMarketingqualifiedleadDate DateTime `json:"hs_lifecyclestage_marketingqualifiedlead_date,omitempty"`

	// The date when a contact's lifecycle stage changed to Opportunity. This is automatically set by HubSpot for each contact.
	HsLifecyclestageOpportunityDate DateTime `json:"hs_lifecyclestage_opportunity_date,omitempty"`

	// The qualification of contacts to sales readiness. It can be set through imports, forms, workflows, and manually on a per contact basis.
	Lifecyclestage Enumeration `json:"lifecyclestage,omitempty"`

	// The date when a contact's lifecycle stage changed to SQL. This is automatically set by HubSpot for each contact.
	HsLifecyclestageSalesqualifiedleadDate DateTime `json:"hs_lifecyclestage_salesqualifiedlead_date,omitempty"`

	// The date that a contact entered the system
	Createdate DateTime `json:"createdate,omitempty"`

	// The date when a contact's lifecycle stage changed to Evangelist. This is automatically set by HubSpot for each contact.
	HsLifecyclestageEvangelistDate DateTime `json:"hs_lifecyclestage_evangelist_date,omitempty"`

	// The date when a contact's lifecycle stage changed to Customer. This is automatically set by HubSpot for each contact.
	HsLifecyclestageCustomerDate DateTime `json:"hs_lifecyclestage_customer_date,omitempty"`

	// The number that shows qualification of contacts to sales readiness. It can be set in HubSpot's Lead Scoring app.
	Hubspotscore Int `json:"hubspotscore,omitempty"`

	// Name of the contact's company. This can be set independently from the name property on the contact's associated company.
	Company string `json:"company,omitempty"`

	// The date when a contact's lifecycle stage changed to Subscriber. This is automatically set by HubSpot for each contact.
	HsLifecyclestageSubscriberDate DateTime `json:"hs_lifecyclestage_subscriber_date,omitempty"`

	// The date when a contact's lifecycle stage changed to Other. This is automatically set by HubSpot for each contact.
	HsLifecyclestageOtherDate DateTime `json:"hs_lifecyclestage_other_date,omitempty"`

	// Associated company website.
	Website string `json:"website,omitempty"`

	// The number of company employees
	Numemployees Enumeration `json:"numemployees,omitempty"`

	// Annual company revenue
	Annualrevenue string `json:"annualrevenue,omitempty"`

	// The Industry a contact is in
	Industry string `json:"industry,omitempty"`

	// HubSpot defined ID of a contact's primary associated company in HubSpot.
	Associatedcompanyid Int `json:"associatedcompanyid,omitempty"`

	// This field is meaningless on its own, and is solely used for triggering dynamic list updates when a company is updated
	Associatedcompanylastupdated Int `json:"associatedcompanylastupdated,omitempty"`

	// The rating of this contact based on their predictive lead score
	HsPredictivecontactscorebucket Enumeration `json:"hs_predictivecontactscorebucket,omitempty"`

	// A score calculated by HubSpot that represents a contact's likelihood to become a customer
	HsPredictivecontactscore Int `json:"hs_predictivecontactscore,omitempty"`
}

type CompanyDefaultProperties struct {
	// Short about-company
	AboutUs string `json:"about_us,omitempty"`

	// Calculation context property providing timestamp for rollup property closedate calculated as EARLIEST_VALUE via values of closedate on object type ObjectTypeId{legacyObjectType=CONTACT}
	ClosedateTimestampEarliestValueA2a17e6e DateTime `json:"closedate_timestamp_earliest_value_a2a17e6e,omitempty"`

	// Number of facebook fans
	Facebookfans Int `json:"facebookfans,omitempty"`

	// Calculation context property providing timestamp for rollup property first_contact_createdate calculated as EARLIEST_VALUE via values of createdate on object type ObjectTypeId{legacyObjectType=CONTACT}
	FirstContactCreatedateTimestampEarliestValue78b50eea DateTime `json:"first_contact_createdate_timestamp_earliest_value_78b50eea,omitempty"`

	// The first conversion date across all contacts associated this company or organization
	FirstConversionDate DateTime `json:"first_conversion_date,omitempty"`

	// Calculation context property providing timestamp for rollup property first_conversion_date calculated as EARLIEST_VALUE via values of first_conversion_date on object type ObjectTypeId{legacyObjectType=CONTACT}
	FirstConversionDateTimestampEarliestValue61f58f2c DateTime `json:"first_conversion_date_timestamp_earliest_value_61f58f2c,omitempty"`

	// The first form submitted across all contacts associated this company or organization
	FirstConversionEventName string `json:"first_conversion_event_name,omitempty"`

	// Calculation context property providing timestamp for rollup property first_conversion_event_name calculated as EARLIEST_VALUE via values of first_conversion_event_name on object type ObjectTypeId{legacyObjectType=CONTACT}
	FirstConversionEventNameTimestampEarliestValue68ddae0a DateTime `json:"first_conversion_event_name_timestamp_earliest_value_68ddae0a,omitempty"`

	// Date the first deal was associated with this company record.
	FirstDealCreatedDate DateTime `json:"first_deal_created_date,omitempty"`

	// The year the company was created. Powered by HubSpot Insights.
	FoundedYear string `json:"founded_year,omitempty"`

	// Additional domains belonging to this company
	HsAdditionalDomains Enumeration `json:"hs_additional_domains,omitempty"`

	// The business units this record is assigned to.
	HsAllAssignedBusinessUnitIds Enumeration `json:"hs_all_assigned_business_unit_ids,omitempty"`

	// The first activity for any contact associated with this company or organization
	HsAnalyticsFirstTimestamp DateTime `json:"hs_analytics_first_timestamp,omitempty"`

	// Calculation context property providing timestamp for rollup property hs_analytics_first_timestamp calculated as EARLIEST_VALUE via values of hs_analytics_first_timestamp on object type ObjectTypeId{legacyObjectType=CONTACT}
	HsAnalyticsFirstTimestampTimestampEarliestValue11e3a63a DateTime `json:"hs_analytics_first_timestamp_timestamp_earliest_value_11e3a63a,omitempty"`

	// The campaign responsible for the first touch creation of the first contact associated with this company
	HsAnalyticsFirstTouchConvertingCampaign string `json:"hs_analytics_first_touch_converting_campaign,omitempty"`

	// Calculation context property providing timestamp for rollup property hs_analytics_first_touch_converting_campaign calculated as EARLIEST_VALUE via values of hs_analytics_first_touch_converting_campaign on object type ObjectTypeId{legacyObjectType=CONTACT}
	HsAnalyticsFirstTouchConvertingCampaignTimestampEarliestValue4757fe10 DateTime `json:"hs_analytics_first_touch_converting_campaign_timestamp_earliest_value_4757fe10,omitempty"`

	// Time of first session across all contacts associated with this company or organization
	HsAnalyticsFirstVisitTimestamp DateTime `json:"hs_analytics_first_visit_timestamp,omitempty"`

	// Calculation context property providing timestamp for rollup property hs_analytics_first_visit_timestamp calculated as EARLIEST_VALUE via values of hs_analytics_first_visit_timestamp on object type ObjectTypeId{legacyObjectType=CONTACT}
	HsAnalyticsFirstVisitTimestampTimestampEarliestValueAccc17ae DateTime `json:"hs_analytics_first_visit_timestamp_timestamp_earliest_value_accc17ae,omitempty"`

	// Time last seen across all contacts associated with this company or organization
	HsAnalyticsLastTimestamp DateTime `json:"hs_analytics_last_timestamp,omitempty"`

	// Calculation context property providing timestamp for rollup property hs_analytics_last_timestamp calculated as LATEST_VALUE via values of hs_analytics_last_timestamp on object type ObjectTypeId{legacyObjectType=CONTACT}
	HsAnalyticsLastTimestampTimestampLatestValue4e16365a DateTime `json:"hs_analytics_last_timestamp_timestamp_latest_value_4e16365a,omitempty"`

	// The campaign responsible for the last touch creation of the first contact associated with this company
	HsAnalyticsLastTouchConvertingCampaign string `json:"hs_analytics_last_touch_converting_campaign,omitempty"`

	// Calculation context property providing timestamp for rollup property hs_analytics_last_touch_converting_campaign calculated as LATEST_VALUE via values of hs_analytics_last_touch_converting_campaign on object type ObjectTypeId{legacyObjectType=CONTACT}
	HsAnalyticsLastTouchConvertingCampaignTimestampLatestValue81a64e30 DateTime `json:"hs_analytics_last_touch_converting_campaign_timestamp_latest_value_81a64e30,omitempty"`

	// Time of the last session attributed to any contacts that are associated with this company record.
	HsAnalyticsLastVisitTimestamp DateTime `json:"hs_analytics_last_visit_timestamp,omitempty"`

	// Calculation context property providing timestamp for rollup property hs_analytics_last_visit_timestamp calculated as LATEST_VALUE via values of hs_analytics_last_visit_timestamp on object type ObjectTypeId{legacyObjectType=CONTACT}
	HsAnalyticsLastVisitTimestampTimestampLatestValue999a0fce DateTime `json:"hs_analytics_last_visit_timestamp_timestamp_latest_value_999a0fce,omitempty"`

	// Source of the last session attributed to any contacts that are associated with this company
	HsAnalyticsLatestSource Enumeration `json:"hs_analytics_latest_source,omitempty"`

	// Additional source details of the last session attributed to any contacts that are associated with this company
	HsAnalyticsLatestSourceData1 string `json:"hs_analytics_latest_source_data_1,omitempty"`

	// Additional source details of the last session attributed to any contacts that are associated with this company
	HsAnalyticsLatestSourceData2 string `json:"hs_analytics_latest_source_data_2,omitempty"`

	// Timestamp of when latest source occurred
	HsAnalyticsLatestSourceTimestamp DateTime `json:"hs_analytics_latest_source_timestamp,omitempty"`

	// Total number of page views across all contacts associated with this company or organization
	HsAnalyticsNumPageViews Int `json:"hs_analytics_num_page_views,omitempty"`

	// Calculation context property providing cardinality for rollup property hs_analytics_num_page_views calculated as SUM via values of hs_analytics_num_page_views on object type ObjectTypeId{legacyObjectType=CONTACT}
	HsAnalyticsNumPageViewsCardinalitySumE46e85b0 Int `json:"hs_analytics_num_page_views_cardinality_sum_e46e85b0,omitempty"`

	// Total number of sessions across all contacts associated with this company or organization
	HsAnalyticsNumVisits Int `json:"hs_analytics_num_visits,omitempty"`

	// Calculation context property providing cardinality for rollup property hs_analytics_num_visits calculated as SUM via values of hs_analytics_num_visits on object type ObjectTypeId{legacyObjectType=CONTACT}
	HsAnalyticsNumVisitsCardinalitySum53d952a6 Int `json:"hs_analytics_num_visits_cardinality_sum_53d952a6,omitempty"`

	// Original source for the contact with the earliest activity for this company or organization
	HsAnalyticsSource Enumeration `json:"hs_analytics_source,omitempty"`

	// Additional information about the original source for the contact with the earliest activity for this company or organization
	HsAnalyticsSourceData1 string `json:"hs_analytics_source_data_1,omitempty"`

	// Calculation context property providing timestamp for rollup property hs_analytics_source_data_1 calculated as EARLIEST_VALUE via values of hs_analytics_source_data_1 on object type ObjectTypeId{legacyObjectType=CONTACT}
	HsAnalyticsSourceData1TimestampEarliestValue9b2f1fa1 DateTime `json:"hs_analytics_source_data_1_timestamp_earliest_value_9b2f1fa1,omitempty"`

	// Additional information about the original source for the contact with the earliest activity for this company or organization
	HsAnalyticsSourceData2 string `json:"hs_analytics_source_data_2,omitempty"`

	// Calculation context property providing timestamp for rollup property hs_analytics_source_data_2 calculated as EARLIEST_VALUE via values of hs_analytics_source_data_2 on object type ObjectTypeId{legacyObjectType=CONTACT}
	HsAnalyticsSourceData2TimestampEarliestValue9b2f9400 DateTime `json:"hs_analytics_source_data_2_timestamp_earliest_value_9b2f9400,omitempty"`

	// Calculation context property providing timestamp for rollup property hs_analytics_source calculated as EARLIEST_VALUE via values of hs_analytics_source on object type ObjectTypeId{legacyObjectType=CONTACT}
	HsAnalyticsSourceTimestampEarliestValue25a3a52c DateTime `json:"hs_analytics_source_timestamp_earliest_value_25a3a52c,omitempty"`

	// The currency code associated with the annual revenue amount
	HsAnnualRevenueCurrencyCode string `json:"hs_annual_revenue_currency_code,omitempty"`

	// The path in the FileManager CDN for this company's avatar override image. Automatically set by HubSpot.
	HsAvatarFilemanagerKey string `json:"hs_avatar_filemanager_key,omitempty"`

	// The user that created this object. This value is automatically set by HubSpot and may not be modified.
	HsCreatedByUserId Int `json:"hs_created_by_user_id,omitempty"`

	// The date and time at which this object was created. This value is automatically set by HubSpot and may not be modified.
	HsCreatedate DateTime `json:"hs_createdate,omitempty"`

	// The date and time when the company entered the 'Customer' stage, 'Lifecycle Stage Pipeline' pipeline
	HsDateEnteredCustomer DateTime `json:"hs_date_entered_customer,omitempty"`

	// The date and time when the company entered the 'Evangelist' stage, 'Lifecycle Stage Pipeline' pipeline
	HsDateEnteredEvangelist DateTime `json:"hs_date_entered_evangelist,omitempty"`

	// The date and time when the company entered the 'Lead' stage, 'Lifecycle Stage Pipeline' pipeline
	HsDateEnteredLead DateTime `json:"hs_date_entered_lead,omitempty"`

	// The date and time when the company entered the 'Marketing Qualified Lead' stage, 'Lifecycle Stage Pipeline' pipeline
	HsDateEnteredMarketingqualifiedlead DateTime `json:"hs_date_entered_marketingqualifiedlead,omitempty"`

	// The date and time when the company entered the 'Opportunity' stage, 'Lifecycle Stage Pipeline' pipeline
	HsDateEnteredOpportunity DateTime `json:"hs_date_entered_opportunity,omitempty"`

	// The date and time when the company entered the 'Other' stage, 'Lifecycle Stage Pipeline' pipeline
	HsDateEnteredOther DateTime `json:"hs_date_entered_other,omitempty"`

	// The date and time when the company entered the 'Sales Qualified Lead' stage, 'Lifecycle Stage Pipeline' pipeline
	HsDateEnteredSalesqualifiedlead DateTime `json:"hs_date_entered_salesqualifiedlead,omitempty"`

	// The date and time when the company entered the 'Subscriber' stage, 'Lifecycle Stage Pipeline' pipeline
	HsDateEnteredSubscriber DateTime `json:"hs_date_entered_subscriber,omitempty"`

	// The date and time when the company exited the 'Customer' stage, 'Lifecycle Stage Pipeline' pipeline
	HsDateExitedCustomer DateTime `json:"hs_date_exited_customer,omitempty"`

	// The date and time when the company exited the 'Evangelist' stage, 'Lifecycle Stage Pipeline' pipeline
	HsDateExitedEvangelist DateTime `json:"hs_date_exited_evangelist,omitempty"`

	// The date and time when the company exited the 'Lead' stage, 'Lifecycle Stage Pipeline' pipeline
	HsDateExitedLead DateTime `json:"hs_date_exited_lead,omitempty"`

	// The date and time when the company exited the 'Marketing Qualified Lead' stage, 'Lifecycle Stage Pipeline' pipeline
	HsDateExitedMarketingqualifiedlead DateTime `json:"hs_date_exited_marketingqualifiedlead,omitempty"`

	// The date and time when the company exited the 'Opportunity' stage, 'Lifecycle Stage Pipeline' pipeline
	HsDateExitedOpportunity DateTime `json:"hs_date_exited_opportunity,omitempty"`

	// The date and time when the company exited the 'Other' stage, 'Lifecycle Stage Pipeline' pipeline
	HsDateExitedOther DateTime `json:"hs_date_exited_other,omitempty"`

	// The date and time when the company exited the 'Sales Qualified Lead' stage, 'Lifecycle Stage Pipeline' pipeline
	HsDateExitedSalesqualifiedlead DateTime `json:"hs_date_exited_salesqualifiedlead,omitempty"`

	// The date and time when the company exited the 'Subscriber' stage, 'Lifecycle Stage Pipeline' pipeline
	HsDateExitedSubscriber DateTime `json:"hs_date_exited_subscriber,omitempty"`

	// How well this company matches your Ideal Customer Profile. Tier 1 means a great fit for your products/services, Tier 3 might be acceptable, but low priority.
	HsIdealCustomerProfile Enumeration `json:"hs_ideal_customer_profile,omitempty"`

	// Identifies whether this company is being marketed and sold to as part of your account-based strategy.
	HsIsTargetAccount Bool `json:"hs_is_target_account,omitempty"`

	// The last date of booked meetings associated with the company
	HsLastBookedMeetingDate DateTime `json:"hs_last_booked_meeting_date,omitempty"`

	// The last date of logged calls associated with the company
	HsLastLoggedCallDate DateTime `json:"hs_last_logged_call_date,omitempty"`

	// The last due date of open tasks associated with the company
	HsLastOpenTaskDate DateTime `json:"hs_last_open_task_date,omitempty"`

	// The date of the last sales activity with the company in seconds.
	HsLastSalesActivityDate DateTime `json:"hs_last_sales_activity_date,omitempty"`

	// The last time a contact engaged with your site or a form, document, meetings link, or tracked email. This doesn't include marketing emails or emails to multiple contacts.
	HsLastSalesActivityTimestamp DateTime `json:"hs_last_sales_activity_timestamp,omitempty"`

	// The type of the last engagement a company performed. This doesn't include marketing emails or emails to multiple contacts.
	HsLastSalesActivityType Enumeration `json:"hs_last_sales_activity_type,omitempty"`

	// Most recent timestamp of any property update for this company. This includes HubSpot internal properties, which can be visible or hidden. This property is updated automatically.
	HsLastmodifieddate DateTime `json:"hs_lastmodifieddate,omitempty"`

	// Latest created date of all associated active Subscriptions
	HsLatestCreatedateOfActiveSubscriptions DateTime `json:"hs_latest_createdate_of_active_subscriptions,omitempty"`

	// The list of Company record IDs that have been merged into this Company. This value is automatically set by HubSpot and may not be modified.
	HsMergedObjectIds Enumeration `json:"hs_merged_object_ids,omitempty"`

	// The number of contacts associated with this company with the role of blocker.
	HsNumBlockers Int `json:"hs_num_blockers,omitempty"`

	// The number of contacts associated with this company with a buying role.
	HsNumContactsWithBuyingRoles Int `json:"hs_num_contacts_with_buying_roles,omitempty"`

	// The number of contacts associated with this company with the role of decision maker.
	HsNumDecisionMakers Int `json:"hs_num_decision_makers,omitempty"`

	// The number of open deals associated with this company.
	HsNumOpenDeals Int `json:"hs_num_open_deals,omitempty"`

	// The unique ID for this record. This value is automatically set by HubSpot and may not be modified.
	HsObjectId Int `json:"hs_object_id,omitempty"`

	// Source (PropertySource) that created this object record
	HsObjectSource string `json:"hs_object_source,omitempty"`

	// First level of detail on how this record was created
	HsObjectSourceDetail1 string `json:"hs_object_source_detail_1,omitempty"`

	// Second level of detail on how this record was created
	HsObjectSourceDetail2 string `json:"hs_object_source_detail_2,omitempty"`

	// Third level of detail on how this record was created
	HsObjectSourceDetail3 string `json:"hs_object_source_detail_3,omitempty"`

	// The sourceId -- further detail -- of the source that created this object record
	HsObjectSourceId string `json:"hs_object_source_id,omitempty"`

	// How this record was created
	HsObjectSourceLabel Enumeration `json:"hs_object_source_label,omitempty"`

	// User ID of the user who initiated creation of this object record
	HsObjectSourceUserId Int `json:"hs_object_source_user_id,omitempty"`

	// The object ID of the current pinned engagement. This value is automatically set by HubSpot and may not be modified.
	HsPinnedEngagementId Int `json:"hs_pinned_engagement_id,omitempty"`

	// The pipeline with which this company is currently associated
	HsPipeline Enumeration `json:"hs_pipeline,omitempty"`

	// The highest probability that a contact associated with this company will become a customer within the next 90 days. This score is based on standard contact properties and behavior.
	HsPredictivecontactscoreV2 Int `json:"hs_predictivecontactscore_v2,omitempty"`

	// Calculation context property providing next_max for rollup property hs_predictivecontactscore_v2 calculated as MAX via values of hs_predictivecontactscore_v2 on object type ObjectTypeId{legacyObjectType=CONTACT}
	HsPredictivecontactscoreV2NextMaxMaxD4e58c1e Int `json:"hs_predictivecontactscore_v2_next_max_max_d4e58c1e,omitempty"`

	// Is the object read only
	HsReadOnly Bool `json:"hs_read_only,omitempty"`

	// The ID of the object from which the data was migrated. This is set automatically during portal data migration.
	HsSourceObjectId Int `json:"hs_source_object_id,omitempty"`

	// The Target Account property is a means to flag high priority companies if you are following an account based strategy
	HsTargetAccount Enumeration `json:"hs_target_account,omitempty"`

	// The probability a company is marked as a target account
	HsTargetAccountProbability Int `json:"hs_target_account_probability,omitempty"`

	// The date when the target account recommendation is snoozed.
	HsTargetAccountRecommendationSnoozeTime DateTime `json:"hs_target_account_recommendation_snooze_time,omitempty"`

	// The state of the target account recommendation
	HsTargetAccountRecommendationState Enumeration `json:"hs_target_account_recommendation_state,omitempty"`

	// The total time in seconds spent by the company in the 'Customer' stage, 'Lifecycle Stage Pipeline' pipeline
	HsTimeInCustomer Int `json:"hs_time_in_customer,omitempty"`

	// The total time in seconds spent by the company in the 'Evangelist' stage, 'Lifecycle Stage Pipeline' pipeline
	HsTimeInEvangelist Int `json:"hs_time_in_evangelist,omitempty"`

	// The total time in seconds spent by the company in the 'Lead' stage, 'Lifecycle Stage Pipeline' pipeline
	HsTimeInLead Int `json:"hs_time_in_lead,omitempty"`

	// The total time in seconds spent by the company in the 'Marketing Qualified Lead' stage, 'Lifecycle Stage Pipeline' pipeline
	HsTimeInMarketingqualifiedlead Int `json:"hs_time_in_marketingqualifiedlead,omitempty"`

	// The total time in seconds spent by the company in the 'Opportunity' stage, 'Lifecycle Stage Pipeline' pipeline
	HsTimeInOpportunity Int `json:"hs_time_in_opportunity,omitempty"`

	// The total time in seconds spent by the company in the 'Other' stage, 'Lifecycle Stage Pipeline' pipeline
	HsTimeInOther Int `json:"hs_time_in_other,omitempty"`

	// The total time in seconds spent by the company in the 'Sales Qualified Lead' stage, 'Lifecycle Stage Pipeline' pipeline
	HsTimeInSalesqualifiedlead Int `json:"hs_time_in_salesqualifiedlead,omitempty"`

	// The total time in seconds spent by the company in the 'Subscriber' stage, 'Lifecycle Stage Pipeline' pipeline
	HsTimeInSubscriber Int `json:"hs_time_in_subscriber,omitempty"`

	// The total value, in your company's currency, of all open deals associated with this company
	HsTotalDealValue Int `json:"hs_total_deal_value,omitempty"`

	// Unique property used for idempotent creates
	HsUniqueCreationKey string `json:"hs_unique_creation_key,omitempty"`

	// The user that last updated this object. This value is automatically set by HubSpot and may not be modified.
	HsUpdatedByUserId Int `json:"hs_updated_by_user_id,omitempty"`

	// The user IDs of all users that have clicked follow within the object to opt-in to getting follow notifications
	HsUserIdsOfAllNotificationFollowers Enumeration `json:"hs_user_ids_of_all_notification_followers,omitempty"`

	// The user IDs of all object owners that have clicked unfollow within the object to opt-out of getting follow notifications
	HsUserIdsOfAllNotificationUnfollowers Enumeration `json:"hs_user_ids_of_all_notification_unfollowers,omitempty"`

	// The user IDs of all owners of this object
	HsUserIdsOfAllOwners Enumeration `json:"hs_user_ids_of_all_owners,omitempty"`

	// Object is part of an import
	HsWasImported Bool `json:"hs_was_imported,omitempty"`

	// The timestamp when an owner was assigned to this company
	HubspotOwnerAssigneddate DateTime `json:"hubspot_owner_assigneddate,omitempty"`

	// Indicates if the company is publicly traded. Powered by HubSpot Insights.
	IsPublic Bool `json:"is_public,omitempty"`

	// The number of contacts associated with this company
	NumAssociatedContacts Int `json:"num_associated_contacts,omitempty"`

	// The number of deals associated with this company
	NumAssociatedDeals Int `json:"num_associated_deals,omitempty"`

	// The number of forms submission for all contacts associated with this company or organization
	NumConversionEvents Int `json:"num_conversion_events,omitempty"`

	// Calculation context property providing cardinality for rollup property num_conversion_events calculated as SUM via values of num_conversion_events on object type ObjectTypeId{legacyObjectType=CONTACT}
	NumConversionEventsCardinalitySumD095f14b Int `json:"num_conversion_events_cardinality_sum_d095f14b,omitempty"`

	// The most recent conversion date across all contacts associated this company or organization
	RecentConversionDate DateTime `json:"recent_conversion_date,omitempty"`

	// Calculation context property providing timestamp for rollup property recent_conversion_date calculated as LATEST_VALUE via values of recent_conversion_date on object type ObjectTypeId{legacyObjectType=CONTACT}
	RecentConversionDateTimestampLatestValue72856da1 DateTime `json:"recent_conversion_date_timestamp_latest_value_72856da1,omitempty"`

	// The last form submitted across all contacts associated this company or organization
	RecentConversionEventName string `json:"recent_conversion_event_name,omitempty"`

	// Calculation context property providing timestamp for rollup property recent_conversion_event_name calculated as LATEST_VALUE via values of recent_conversion_event_name on object type ObjectTypeId{legacyObjectType=CONTACT}
	RecentConversionEventNameTimestampLatestValue66c820bf DateTime `json:"recent_conversion_event_name_timestamp_latest_value_66c820bf,omitempty"`

	// Amount of last closed won deal associated with this company. Set automatically.
	RecentDealAmount Int `json:"recent_deal_amount,omitempty"`

	// Date of the last "closed won" deal associated with this company record.
	RecentDealCloseDate DateTime `json:"recent_deal_close_date,omitempty"`

	// Time zone where the company or organization is located. Powered by HubSpot Insights.
	Timezone string `json:"timezone,omitempty"`

	// The total amount of money raised by the company. Powered by HubSpot Insights.
	TotalMoneyRaised string `json:"total_money_raised,omitempty"`

	// The total amount of closed won deals
	TotalRevenue Int `json:"total_revenue,omitempty"`

	// The name of the company or organization. Powered by HubSpot Insights.
	Name string `json:"name,omitempty"`

	// HubSpot owner email for this company or organization
	Owneremail string `json:"owneremail,omitempty"`

	// The main twitter account of the company or organization
	Twitterhandle string `json:"twitterhandle,omitempty"`

	// HubSpot owner name for this company or organization
	Ownername string `json:"ownername,omitempty"`

	// Company primary phone number. Powered by HubSpot Insights.
	Phone string `json:"phone,omitempty"`

	// The Twitter bio of the company or organization
	Twitterbio string `json:"twitterbio,omitempty"`

	// The number of Twitter followers of the company or organization
	Twitterfollowers Int `json:"twitterfollowers,omitempty"`

	// Street address of the company or organization, including unit number. Powered by HubSpot Insights.
	Address string `json:"address,omitempty"`

	// Additional address of the company or organization. Powered by HubSpot Insights.
	Address2 string `json:"address2,omitempty"`

	// The URL of the Facebook company page for the company or organization
	FacebookCompanyPage string `json:"facebook_company_page,omitempty"`

	// City where the company is located. Powered by HubSpot Insights.
	City string `json:"city,omitempty"`

	// The URL of the LinkedIn company page for the company or organization
	LinkedinCompanyPage string `json:"linkedin_company_page,omitempty"`

	// The LinkedIn bio for the company or organization
	Linkedinbio string `json:"linkedinbio,omitempty"`

	// State or region in which the company or organization is located. Powered by HubSpot Insights.
	State string `json:"state,omitempty"`

	// The URL of the Google Plus page for the company or organization
	GoogleplusPage string `json:"googleplus_page,omitempty"`

	// Last Meeting Booked for any contact associated with this Company/Organization.
	EngagementsLastMeetingBooked DateTime `json:"engagements_last_meeting_booked,omitempty"`

	// This UTM parameter shows which marketing campaign (e.g. a specific email) referred an associated contact to the meetings tool for their most recent booking. This property is only populated when you add tracking parameters to your meeting link.
	EngagementsLastMeetingBookedCampaign string `json:"engagements_last_meeting_booked_campaign,omitempty"`

	// This UTM parameter shows which channel (e.g. email) referred an associated contact to the meetings tool for their most recent booking. This property is only populated when you add tracking parameters to your meeting link.
	EngagementsLastMeetingBookedMedium string `json:"engagements_last_meeting_booked_medium,omitempty"`

	// This UTM parameter shows which site (e.g. Twitter) referred an associated contact to the meetings tool for their most recent booking. This property is only populated when you add tracking parameters to your meeting link.
	EngagementsLastMeetingBookedSource string `json:"engagements_last_meeting_booked_source,omitempty"`

	// The date of the most recent meeting (past or upcoming) logged for, scheduled with, or booked by a contact associated with this company.
	HsLatestMeetingActivity DateTime `json:"hs_latest_meeting_activity,omitempty"`

	// The last time a tracked sales email was replied to by this company
	HsSalesEmailLastReplied DateTime `json:"hs_sales_email_last_replied,omitempty"`

	// The owner of the company
	HubspotOwnerId Enumeration `json:"hubspot_owner_id,omitempty"`

	// The last time a call, chat conversation, LinkedIn message, postal mail, meeting, sales email, SMS, or WhatsApp message was logged for a company. This is set automatically by HubSpot based on user actions in the company record.
	NotesLastContacted DateTime `json:"notes_last_contacted,omitempty"`

	// The last time a call, chat conversation, LinkedIn message, postal mail, meeting, note, sales email, SMS, or WhatsApp message was logged for a company. This is set automatically by HubSpot based on user actions in the company record.
	NotesLastUpdated DateTime `json:"notes_last_updated,omitempty"`

	// Date of the next upcoming scheduled sales activity for this company record.
	NotesNextActivityDate DateTime `json:"notes_next_activity_date,omitempty"`

	// The number of times a call, chat conversation, LinkedIn message, postal mail, meeting, sales email, SMS, or WhatsApp message was logged for a company record. This is set automatically by HubSpot based on user actions in the company record.
	NumContactedNotes Int `json:"num_contacted_notes,omitempty"`

	// The number of times a call, chat conversation, LinkedIn message, postal mail, meeting, note, sales email, SMS, task, or WhatsApp message was logged for a company record. This is set automatically by HubSpot based on user actions in the company record.
	NumNotes Int `json:"num_notes,omitempty"`

	// Postal or zip code of the company or organization. Powered by HubSpot Insights.
	Zip string `json:"zip,omitempty"`

	// Country in which the company or organization is located. Powered by HubSpot Insights.
	Country string `json:"country,omitempty"`

	// The team of the owner of the company.
	HubspotTeamId Enumeration `json:"hubspot_team_id,omitempty"`

	// The value of all owner referencing properties for this object, both default and custom
	HsAllOwnerIds Enumeration `json:"hs_all_owner_ids,omitempty"`

	// The main website of the company or organization. This property is used to identify unique companies. Powered by HubSpot Insights.
	Website string `json:"website,omitempty"`

	// The domain name of the company or organization
	Domain string `json:"domain,omitempty"`

	// The team ids corresponding to all owner referencing properties for this object, both default and custom
	HsAllTeamIds Enumeration `json:"hs_all_team_ids,omitempty"`

	// The team ids, including up the team hierarchy, corresponding to all owner referencing properties for this object, both default and custom
	HsAllAccessibleTeamIds Enumeration `json:"hs_all_accessible_team_ids,omitempty"`

	// The total number of employees who work for the company or organization. Powered by HubSpot Insights.
	Numberofemployees Int `json:"numberofemployees,omitempty"`

	// The type of business the company performs. By default, this property has approximately 150 pre-defined options to select from. While these options cannot be deleted as they used by HubSpot Insights, you can add new custom options to meet your needs.
	Industry Enumeration `json:"industry,omitempty"`

	// The actual or estimated annual revenue of the company. Powered by HubSpot Insights.
	Annualrevenue Int `json:"annualrevenue,omitempty"`

	// The qualification of companies to sales readiness throughout the buying journey
	Lifecyclestage Enumeration `json:"lifecyclestage,omitempty"`

	// The company's sales, prospecting or outreach status
	HsLeadStatus Enumeration `json:"hs_lead_status,omitempty"`

	// The parent company of this company
	HsParentCompanyId Int `json:"hs_parent_company_id,omitempty"`

	// The optional classification of this company record - prospect, partner, etc.
	Type Enumeration `json:"type,omitempty"`

	// A short statement about the company's mission and goals. Powered by HubSpot Insights.
	Description string `json:"description,omitempty"`

	// The number of child companies of this company
	HsNumChildCompanies Int `json:"hs_num_child_companies,omitempty"`

	// The sales and marketing score for the company or organization
	Hubspotscore Int `json:"hubspotscore,omitempty"`

	// The date the company or organization was added to the database
	Createdate DateTime `json:"createdate,omitempty"`

	// The date the company or organization was closed as a customer
	Closedate DateTime `json:"closedate,omitempty"`

	// The date that the first contact from this company entered the system, which could pre-date the company's create date
	FirstContactCreatedate DateTime `json:"first_contact_createdate,omitempty"`

	// The number of days between when the company record was created and when they closed as a customer.
	DaysToClose Int `json:"days_to_close,omitempty"`

	// The web technologies used by the company or organization. Powered by HubSpot Insights.
	WebTechnologies Enumeration `json:"web_technologies,omitempty"`
}

type DealDefaultProperties struct {
	// The amount of the deal, using the exchange rate, in your company's currency
	AmountInHomeCurrency Int `json:"amount_in_home_currency,omitempty"`

	// The number of days the deal took to close
	DaysToClose Int `json:"days_to_close,omitempty"`

	// Currency code for the deal.
	DealCurrencyCode Enumeration `json:"deal_currency_code,omitempty"`

	// The annual contract value (ACV) of this deal.
	HsAcv Int `json:"hs_acv,omitempty"`

	// The business units this record is assigned to.
	HsAllAssignedBusinessUnitIds Enumeration `json:"hs_all_assigned_business_unit_ids,omitempty"`

	// Owner ids of the users involved in closing the deal
	HsAllCollaboratorOwnerIds Enumeration `json:"hs_all_collaborator_owner_ids,omitempty"`

	// The owner ids of all associated Deal Splits. This property is set automatically by HubSpot.
	HsAllDealSplitOwnerIds Enumeration `json:"hs_all_deal_split_owner_ids,omitempty"`

	// Source for the contact either directly or indirectly associated with the last session activity for this deal
	HsAnalyticsLatestSource Enumeration `json:"hs_analytics_latest_source,omitempty"`

	// Source for the company with an associated contact with the last session activity for this deal
	HsAnalyticsLatestSourceCompany Enumeration `json:"hs_analytics_latest_source_company,omitempty"`

	// Source for the directly associated contact with the last session activity for this deal
	HsAnalyticsLatestSourceContact Enumeration `json:"hs_analytics_latest_source_contact,omitempty"`

	// Additional source details of the last session attributed to any contacts that are directly or indirectly associated with this deal
	HsAnalyticsLatestSourceData1 string `json:"hs_analytics_latest_source_data_1,omitempty"`

	// Additional source details of the last session attributed to any contacts that are indirectly associated with this deal (via a company association)
	HsAnalyticsLatestSourceData1Company string `json:"hs_analytics_latest_source_data_1_company,omitempty"`

	// Additional source details of the last session attributed to any contacts that are directly associated with this deal
	HsAnalyticsLatestSourceData1Contact string `json:"hs_analytics_latest_source_data_1_contact,omitempty"`

	// Additional source details of the last session attributed to any contacts that are directly or indirectly associated with this deal
	HsAnalyticsLatestSourceData2 string `json:"hs_analytics_latest_source_data_2,omitempty"`

	// Additional source details of the last session attributed to any contacts that are indirectly associated with this deal
	HsAnalyticsLatestSourceData2Company string `json:"hs_analytics_latest_source_data_2_company,omitempty"`

	// Additional source details of the last session attributed to any contacts that are directly associated with this deal
	HsAnalyticsLatestSourceData2Contact string `json:"hs_analytics_latest_source_data_2_contact,omitempty"`

	// Timestamp of when latest source occurred for either a directly or indirectly associated contact
	HsAnalyticsLatestSourceTimestamp DateTime `json:"hs_analytics_latest_source_timestamp,omitempty"`

	// Timestamp of when latest source occurred for an indirectly associated contact
	HsAnalyticsLatestSourceTimestampCompany DateTime `json:"hs_analytics_latest_source_timestamp_company,omitempty"`

	// Timestamp of when latest source occurred for a directly associated contact
	HsAnalyticsLatestSourceTimestampContact DateTime `json:"hs_analytics_latest_source_timestamp_contact,omitempty"`

	// Original source for the contact with the earliest activity for this deal.
	HsAnalyticsSource Enumeration `json:"hs_analytics_source,omitempty"`

	// Additional information about the original source for the associated contact, or associated company if there is no contact, with the oldest value for the Time first seen property.
	HsAnalyticsSourceData1 string `json:"hs_analytics_source_data_1,omitempty"`

	// Additional information about the original source for the associated contact, or associated company if there is no contact, with the oldest value for the Time first seen property.
	HsAnalyticsSourceData2 string `json:"hs_analytics_source_data_2,omitempty"`

	// The annual recurring revenue (ARR) of this deal.
	HsArr Int `json:"hs_arr,omitempty"`

	// The marketing campaign the deal is associated with
	HsCampaign string `json:"hs_campaign,omitempty"`

	// Returns the amount if the deal is closed. Else, returns 0.
	HsClosedAmount Int `json:"hs_closed_amount,omitempty"`

	// Returns the amount in home currency if the deal is closed. Else, returns 0.
	HsClosedAmountInHomeCurrency Int `json:"hs_closed_amount_in_home_currency,omitempty"`

	// This property is 1 if the deal is closed won, otherwise 0.
	HsClosedWonCount Int `json:"hs_closed_won_count,omitempty"`

	// Returns closedate if this deal is closed won
	HsClosedWonDate DateTime `json:"hs_closed_won_date,omitempty"`

	// The user that created this object. This value is automatically set by HubSpot and may not be modified.
	HsCreatedByUserId Int `json:"hs_created_by_user_id,omitempty"`

	// The date and time when the deal entered the 'Appointment Scheduled' stage, 'Sales Pipeline' pipeline
	HsDateEnteredAppointmentscheduled DateTime `json:"hs_date_entered_appointmentscheduled,omitempty"`

	// The date and time when the deal entered the 'Closed Lost' stage, 'Sales Pipeline' pipeline
	HsDateEnteredClosedlost DateTime `json:"hs_date_entered_closedlost,omitempty"`

	// The date and time when the deal entered the 'Closed Won' stage, 'Sales Pipeline' pipeline
	HsDateEnteredClosedwon DateTime `json:"hs_date_entered_closedwon,omitempty"`

	// The date and time when the deal entered the 'Contract Sent' stage, 'Sales Pipeline' pipeline
	HsDateEnteredContractsent DateTime `json:"hs_date_entered_contractsent,omitempty"`

	// The date and time when the deal entered the 'Decision Maker Bought-In' stage, 'Sales Pipeline' pipeline
	HsDateEnteredDecisionmakerboughtin DateTime `json:"hs_date_entered_decisionmakerboughtin,omitempty"`

	// The date and time when the deal entered the 'Presentation Scheduled' stage, 'Sales Pipeline' pipeline
	HsDateEnteredPresentationscheduled DateTime `json:"hs_date_entered_presentationscheduled,omitempty"`

	// The date and time when the deal entered the 'Qualified To Buy' stage, 'Sales Pipeline' pipeline
	HsDateEnteredQualifiedtobuy DateTime `json:"hs_date_entered_qualifiedtobuy,omitempty"`

	// The date and time when the deal exited the 'Appointment Scheduled' stage, 'Sales Pipeline' pipeline
	HsDateExitedAppointmentscheduled DateTime `json:"hs_date_exited_appointmentscheduled,omitempty"`

	// The date and time when the deal exited the 'Closed Lost' stage, 'Sales Pipeline' pipeline
	HsDateExitedClosedlost DateTime `json:"hs_date_exited_closedlost,omitempty"`

	// The date and time when the deal exited the 'Closed Won' stage, 'Sales Pipeline' pipeline
	HsDateExitedClosedwon DateTime `json:"hs_date_exited_closedwon,omitempty"`

	// The date and time when the deal exited the 'Contract Sent' stage, 'Sales Pipeline' pipeline
	HsDateExitedContractsent DateTime `json:"hs_date_exited_contractsent,omitempty"`

	// The date and time when the deal exited the 'Decision Maker Bought-In' stage, 'Sales Pipeline' pipeline
	HsDateExitedDecisionmakerboughtin DateTime `json:"hs_date_exited_decisionmakerboughtin,omitempty"`

	// The date and time when the deal exited the 'Presentation Scheduled' stage, 'Sales Pipeline' pipeline
	HsDateExitedPresentationscheduled DateTime `json:"hs_date_exited_presentationscheduled,omitempty"`

	// The date and time when the deal exited the 'Qualified To Buy' stage, 'Sales Pipeline' pipeline
	HsDateExitedQualifiedtobuy DateTime `json:"hs_date_exited_qualifiedtobuy,omitempty"`

	// The number of days the deal took to close, without rounding
	HsDaysToCloseRaw Int `json:"hs_days_to_close_raw,omitempty"`

	// Specifies how deal amount should be calculated from line items
	HsDealAmountCalculationPreference Enumeration `json:"hs_deal_amount_calculation_preference,omitempty"`

	// The predictive deal score calculated by Hubspot AI to score the deal health
	HsDealScore Int `json:"hs_deal_score,omitempty"`

	// The probability a deal will close. This defaults to the deal stage probability setting.
	HsDealStageProbability Int `json:"hs_deal_stage_probability,omitempty"`

	// Fall back property for calculating the deal stage when no customer override exist.  Probability between 0 and 1 of deal stage. Defaults to 0 for unknown deal stages.
	HsDealStageProbabilityShadow Int `json:"hs_deal_stage_probability_shadow,omitempty"`

	// This is the exchange rate used to convert the deal amount into your company currency.
	HsExchangeRate Int `json:"hs_exchange_rate,omitempty"`

	// The custom forecasted deal value calculated by multiplying the forecast probability and deal amount in your company’s currency.
	HsForecastAmount Int `json:"hs_forecast_amount,omitempty"`

	// The custom percent probability a deal will close.
	HsForecastProbability Int `json:"hs_forecast_probability,omitempty"`

	// Indicates if the current deal is an active shared deal. It is set automatically based on the value of hs_num_associated_active_deal_registrations.
	HsIsActiveSharedDeal Bool `json:"hs_is_active_shared_deal,omitempty"`

	// True if the deal was won or lost.
	HsIsClosed Bool `json:"hs_is_closed,omitempty"`

	// True if the deal is in the closed won state, false otherwise
	HsIsClosedWon Bool `json:"hs_is_closed_won,omitempty"`

	// Indicates if the deal is split between multiple users.
	HsIsDealSplit Bool `json:"hs_is_deal_split,omitempty"`

	// This property is 1 if the deal is not closed won or closed lost, otherwise 0
	HsIsOpenCount Int `json:"hs_is_open_count,omitempty"`

	// Most recent timestamp of any property update for this deal. This includes HubSpot internal properties, which can be visible or hidden. This property is updated automatically.
	HsLastmodifieddate DateTime `json:"hs_lastmodifieddate,omitempty"`

	// Hubspot predicted likelihood between 0 and 1 of the deal to close by the close date.
	HsLikelihoodToClose Int `json:"hs_likelihood_to_close,omitempty"`

	// For internal HubSpot Application use only. Global term for the discount percentage applied.
	HsLineItemGlobalTermHsDiscountPercentage string `json:"hs_line_item_global_term_hs_discount_percentage,omitempty"`

	// For internal HubSpot Application use only. Indicates if the Global term for the discount percentage is enabled.
	HsLineItemGlobalTermHsDiscountPercentageEnabled Bool `json:"hs_line_item_global_term_hs_discount_percentage_enabled,omitempty"`

	// For internal HubSpot Application use only. Global term for product recurring billing duration.
	HsLineItemGlobalTermHsRecurringBillingPeriod string `json:"hs_line_item_global_term_hs_recurring_billing_period,omitempty"`

	// For internal HubSpot Application use only. Indicates if the Global term for product recurring billing duration is enabled.
	HsLineItemGlobalTermHsRecurringBillingPeriodEnabled Bool `json:"hs_line_item_global_term_hs_recurring_billing_period_enabled,omitempty"`

	// For internal HubSpot Application use only. Global term for recurring billing start date for a line item.
	HsLineItemGlobalTermHsRecurringBillingStartDate string `json:"hs_line_item_global_term_hs_recurring_billing_start_date,omitempty"`

	// For internal HubSpot Application use only. Indicates if the Global term for recurring billing start date for a line item is enabled.
	HsLineItemGlobalTermHsRecurringBillingStartDateEnabled Bool `json:"hs_line_item_global_term_hs_recurring_billing_start_date_enabled,omitempty"`

	// For internal HubSpot Application use only. Global term for how frequently the product is billed.
	HsLineItemGlobalTermRecurringbillingfrequency string `json:"hs_line_item_global_term_recurringbillingfrequency,omitempty"`

	// For internal HubSpot Application use only. Indicates if the Global term for how frequently the product is billed is enabled
	HsLineItemGlobalTermRecurringbillingfrequencyEnabled Bool `json:"hs_line_item_global_term_recurringbillingfrequency_enabled,omitempty"`

	// The likelihood a deal will close. This property is used for manual forecasting your deals.
	HsManualForecastCategory Enumeration `json:"hs_manual_forecast_category,omitempty"`

	// The list of Deal record IDs that have been merged into this Deal. This value is automatically set by HubSpot and may not be modified.
	HsMergedObjectIds Enumeration `json:"hs_merged_object_ids,omitempty"`

	// The monthly recurring revenue (MRR) of this deal.
	HsMrr Int `json:"hs_mrr,omitempty"`

	// A short description of the next step for the deal
	HsNextStep string `json:"hs_next_step,omitempty"`

	// The number of active deal registrations associated with this deal. This property is set automatically by HubSpot.
	HsNumAssociatedActiveDealRegistrations Int `json:"hs_num_associated_active_deal_registrations,omitempty"`

	// The number of deal registrations associated with this deal. This property is set automatically by HubSpot.
	HsNumAssociatedDealRegistrations Int `json:"hs_num_associated_deal_registrations,omitempty"`

	// The number of deal splits associated with this deal. This property is set automatically by HubSpot.
	HsNumAssociatedDealSplits Int `json:"hs_num_associated_deal_splits,omitempty"`

	// The number of line items associated with this deal
	HsNumOfAssociatedLineItems Int `json:"hs_num_of_associated_line_items,omitempty"`

	// The number of target account companies associated with this deal. This property is set automatically by HubSpot.
	HsNumTargetAccounts Int `json:"hs_num_target_accounts,omitempty"`

	// The unique ID for this record. This value is automatically set by HubSpot and may not be modified.
	HsObjectId Int `json:"hs_object_id,omitempty"`

	// Source (PropertySource) that created this object record
	HsObjectSource string `json:"hs_object_source,omitempty"`

	// First level of detail on how this record was created
	HsObjectSourceDetail1 string `json:"hs_object_source_detail_1,omitempty"`

	// Second level of detail on how this record was created
	HsObjectSourceDetail2 string `json:"hs_object_source_detail_2,omitempty"`

	// Third level of detail on how this record was created
	HsObjectSourceDetail3 string `json:"hs_object_source_detail_3,omitempty"`

	// The sourceId -- further detail -- of the source that created this object record
	HsObjectSourceId string `json:"hs_object_source_id,omitempty"`

	// How this record was created
	HsObjectSourceLabel Enumeration `json:"hs_object_source_label,omitempty"`

	// User ID of the user who initiated creation of this object record
	HsObjectSourceUserId Int `json:"hs_object_source_user_id,omitempty"`

	// The object ID of the current pinned engagement. This value is automatically set by HubSpot and may not be modified.
	HsPinnedEngagementId Int `json:"hs_pinned_engagement_id,omitempty"`

	// Returns the multiplication of the deal amount times the predicted likelihood of the deal to close by the close date.
	HsPredictedAmount Int `json:"hs_predicted_amount,omitempty"`

	// Returns the multiplication of the deal amount in your company's currency times the predicted likelihood of the deal to close by the close date.
	HsPredictedAmountInHomeCurrency Int `json:"hs_predicted_amount_in_home_currency,omitempty"`

	//
	HsPriority Enumeration `json:"hs_priority,omitempty"`

	// Returns the multiplication of the amount times the probability of the deal closing.
	HsProjectedAmount Int `json:"hs_projected_amount,omitempty"`

	// Returns the multiplication of the amount in home currency times the probability of the deal closing.
	HsProjectedAmountInHomeCurrency Int `json:"hs_projected_amount_in_home_currency,omitempty"`

	// Is the object read only
	HsReadOnly Bool `json:"hs_read_only,omitempty"`

	// The ID of the object from which the data was migrated. This is set automatically during portal data migration.
	HsSourceObjectId Int `json:"hs_source_object_id,omitempty"`

	// List of tag ids applicable to a deal. This property is set automatically by HubSpot.
	HsTagIds Enumeration `json:"hs_tag_ids,omitempty"`

	// The total contract value (TCV) of this deal.
	HsTcv Int `json:"hs_tcv,omitempty"`

	// The total time in seconds spent by the deal in the 'Appointment Scheduled' stage, 'Sales Pipeline' pipeline
	HsTimeInAppointmentscheduled Int `json:"hs_time_in_appointmentscheduled,omitempty"`

	// The total time in seconds spent by the deal in the 'Closed Lost' stage, 'Sales Pipeline' pipeline
	HsTimeInClosedlost Int `json:"hs_time_in_closedlost,omitempty"`

	// The total time in seconds spent by the deal in the 'Closed Won' stage, 'Sales Pipeline' pipeline
	HsTimeInClosedwon Int `json:"hs_time_in_closedwon,omitempty"`

	// The total time in seconds spent by the deal in the 'Contract Sent' stage, 'Sales Pipeline' pipeline
	HsTimeInContractsent Int `json:"hs_time_in_contractsent,omitempty"`

	// The total time in seconds spent by the deal in the 'Decision Maker Bought-In' stage, 'Sales Pipeline' pipeline
	HsTimeInDecisionmakerboughtin Int `json:"hs_time_in_decisionmakerboughtin,omitempty"`

	// The total time in seconds spent by the deal in the 'Presentation Scheduled' stage, 'Sales Pipeline' pipeline
	HsTimeInPresentationscheduled Int `json:"hs_time_in_presentationscheduled,omitempty"`

	// The total time in seconds spent by the deal in the 'Qualified To Buy' stage, 'Sales Pipeline' pipeline
	HsTimeInQualifiedtobuy Int `json:"hs_time_in_qualifiedtobuy,omitempty"`

	// Unique property used for idempotent creates
	HsUniqueCreationKey string `json:"hs_unique_creation_key,omitempty"`

	// The user that last updated this object. This value is automatically set by HubSpot and may not be modified.
	HsUpdatedByUserId Int `json:"hs_updated_by_user_id,omitempty"`

	// The user IDs of all users that have clicked follow within the object to opt-in to getting follow notifications
	HsUserIdsOfAllNotificationFollowers Enumeration `json:"hs_user_ids_of_all_notification_followers,omitempty"`

	// The user IDs of all object owners that have clicked unfollow within the object to opt-out of getting follow notifications
	HsUserIdsOfAllNotificationUnfollowers Enumeration `json:"hs_user_ids_of_all_notification_unfollowers,omitempty"`

	// The user IDs of all owners of this object
	HsUserIdsOfAllOwners Enumeration `json:"hs_user_ids_of_all_owners,omitempty"`

	// The cumulative time in seconds spent by the deal  in the 'Appointment Scheduled' stage, 'Sales Pipeline' pipeline
	HsV2CumulativeTimeInAppointmentscheduled Int `json:"hs_v2_cumulative_time_in_appointmentscheduled,omitempty"`

	// The cumulative time in seconds spent by the deal  in the 'Closed Lost' stage, 'Sales Pipeline' pipeline
	HsV2CumulativeTimeInClosedlost Int `json:"hs_v2_cumulative_time_in_closedlost,omitempty"`

	// The cumulative time in seconds spent by the deal  in the 'Closed Won' stage, 'Sales Pipeline' pipeline
	HsV2CumulativeTimeInClosedwon Int `json:"hs_v2_cumulative_time_in_closedwon,omitempty"`

	// The cumulative time in seconds spent by the deal  in the 'Contract Sent' stage, 'Sales Pipeline' pipeline
	HsV2CumulativeTimeInContractsent Int `json:"hs_v2_cumulative_time_in_contractsent,omitempty"`

	// The cumulative time in seconds spent by the deal  in the 'Decision Maker Bought-In' stage, 'Sales Pipeline' pipeline
	HsV2CumulativeTimeInDecisionmakerboughtin Int `json:"hs_v2_cumulative_time_in_decisionmakerboughtin,omitempty"`

	// The cumulative time in seconds spent by the deal  in the 'Presentation Scheduled' stage, 'Sales Pipeline' pipeline
	HsV2CumulativeTimeInPresentationscheduled Int `json:"hs_v2_cumulative_time_in_presentationscheduled,omitempty"`

	// The cumulative time in seconds spent by the deal  in the 'Qualified To Buy' stage, 'Sales Pipeline' pipeline
	HsV2CumulativeTimeInQualifiedtobuy Int `json:"hs_v2_cumulative_time_in_qualifiedtobuy,omitempty"`

	// The date and time when the deal entered the 'Appointment Scheduled' stage, 'Sales Pipeline' pipeline
	HsV2DateEnteredAppointmentscheduled DateTime `json:"hs_v2_date_entered_appointmentscheduled,omitempty"`

	// The date and time when the deal entered the 'Closed Lost' stage, 'Sales Pipeline' pipeline
	HsV2DateEnteredClosedlost DateTime `json:"hs_v2_date_entered_closedlost,omitempty"`

	// The date and time when the deal entered the 'Closed Won' stage, 'Sales Pipeline' pipeline
	HsV2DateEnteredClosedwon DateTime `json:"hs_v2_date_entered_closedwon,omitempty"`

	// The date and time when the deal entered the 'Contract Sent' stage, 'Sales Pipeline' pipeline
	HsV2DateEnteredContractsent DateTime `json:"hs_v2_date_entered_contractsent,omitempty"`

	// The date and time when the deal entered the 'Decision Maker Bought-In' stage, 'Sales Pipeline' pipeline
	HsV2DateEnteredDecisionmakerboughtin DateTime `json:"hs_v2_date_entered_decisionmakerboughtin,omitempty"`

	// The date and time when the deal entered the 'Presentation Scheduled' stage, 'Sales Pipeline' pipeline
	HsV2DateEnteredPresentationscheduled DateTime `json:"hs_v2_date_entered_presentationscheduled,omitempty"`

	// The date and time when the deal entered the 'Qualified To Buy' stage, 'Sales Pipeline' pipeline
	HsV2DateEnteredQualifiedtobuy DateTime `json:"hs_v2_date_entered_qualifiedtobuy,omitempty"`

	// The date and time when the deal exited the 'Appointment Scheduled' stage, 'Sales Pipeline' pipeline
	HsV2DateExitedAppointmentscheduled DateTime `json:"hs_v2_date_exited_appointmentscheduled,omitempty"`

	// The date and time when the deal exited the 'Closed Lost' stage, 'Sales Pipeline' pipeline
	HsV2DateExitedClosedlost DateTime `json:"hs_v2_date_exited_closedlost,omitempty"`

	// The date and time when the deal exited the 'Closed Won' stage, 'Sales Pipeline' pipeline
	HsV2DateExitedClosedwon DateTime `json:"hs_v2_date_exited_closedwon,omitempty"`

	// The date and time when the deal exited the 'Contract Sent' stage, 'Sales Pipeline' pipeline
	HsV2DateExitedContractsent DateTime `json:"hs_v2_date_exited_contractsent,omitempty"`

	// The date and time when the deal exited the 'Decision Maker Bought-In' stage, 'Sales Pipeline' pipeline
	HsV2DateExitedDecisionmakerboughtin DateTime `json:"hs_v2_date_exited_decisionmakerboughtin,omitempty"`

	// The date and time when the deal exited the 'Presentation Scheduled' stage, 'Sales Pipeline' pipeline
	HsV2DateExitedPresentationscheduled DateTime `json:"hs_v2_date_exited_presentationscheduled,omitempty"`

	// The date and time when the deal exited the 'Qualified To Buy' stage, 'Sales Pipeline' pipeline
	HsV2DateExitedQualifiedtobuy DateTime `json:"hs_v2_date_exited_qualifiedtobuy,omitempty"`

	// The total time in seconds spent by the deal in the 'Appointment Scheduled' stage, 'Sales Pipeline' pipeline since it last entered this stage
	HsV2LatestTimeInAppointmentscheduled Int `json:"hs_v2_latest_time_in_appointmentscheduled,omitempty"`

	// The total time in seconds spent by the deal in the 'Closed Lost' stage, 'Sales Pipeline' pipeline since it last entered this stage
	HsV2LatestTimeInClosedlost Int `json:"hs_v2_latest_time_in_closedlost,omitempty"`

	// The total time in seconds spent by the deal in the 'Closed Won' stage, 'Sales Pipeline' pipeline since it last entered this stage
	HsV2LatestTimeInClosedwon Int `json:"hs_v2_latest_time_in_closedwon,omitempty"`

	// The total time in seconds spent by the deal in the 'Contract Sent' stage, 'Sales Pipeline' pipeline since it last entered this stage
	HsV2LatestTimeInContractsent Int `json:"hs_v2_latest_time_in_contractsent,omitempty"`

	// The total time in seconds spent by the deal in the 'Decision Maker Bought-In' stage, 'Sales Pipeline' pipeline since it last entered this stage
	HsV2LatestTimeInDecisionmakerboughtin Int `json:"hs_v2_latest_time_in_decisionmakerboughtin,omitempty"`

	// The total time in seconds spent by the deal in the 'Presentation Scheduled' stage, 'Sales Pipeline' pipeline since it last entered this stage
	HsV2LatestTimeInPresentationscheduled Int `json:"hs_v2_latest_time_in_presentationscheduled,omitempty"`

	// The total time in seconds spent by the deal in the 'Qualified To Buy' stage, 'Sales Pipeline' pipeline since it last entered this stage
	HsV2LatestTimeInQualifiedtobuy Int `json:"hs_v2_latest_time_in_qualifiedtobuy,omitempty"`

	// Object is part of an import
	HsWasImported Bool `json:"hs_was_imported,omitempty"`

	// The date the most recent deal owner was assigned to a deal. This is updated automatically by HubSpot.
	HubspotOwnerAssigneddate DateTime `json:"hubspot_owner_assigneddate,omitempty"`

	// The name given to this deal.
	Dealname string `json:"dealname,omitempty"`

	// The total amount of the deal
	Amount Int `json:"amount,omitempty"`

	// The stage of the deal. Deal stages allow you to categorize and track the progress of the deals that you are working on.
	Dealstage Enumeration `json:"dealstage,omitempty"`

	// The pipeline the deal is in. This determines which stages are options for the deal.
	Pipeline Enumeration `json:"pipeline,omitempty"`

	// Date the deal was closed. This property is set automatically by HubSpot.
	Closedate DateTime `json:"closedate,omitempty"`

	// Date the deal was created. This property is set automatically by HubSpot.
	Createdate DateTime `json:"createdate,omitempty"`

	// The date of the most recent meeting an associated contact has booked through the meetings tool.
	EngagementsLastMeetingBooked DateTime `json:"engagements_last_meeting_booked,omitempty"`

	// This UTM parameter shows which marketing campaign (e.g. a specific email) referred an associated contact to the meetings tool for their most recent booking. This property is only populated when you add tracking parameters to your meeting link.
	EngagementsLastMeetingBookedCampaign string `json:"engagements_last_meeting_booked_campaign,omitempty"`

	// This UTM parameter shows which channel (e.g. email) referred an associated contact to the meetings tool for their most recent booking. This property is only populated when you add tracking parameters to your meeting link.
	EngagementsLastMeetingBookedMedium string `json:"engagements_last_meeting_booked_medium,omitempty"`

	// This UTM parameter shows which site (e.g. Twitter) referred an associated contact to the meetings tool for their most recent booking. This property is only populated when you add tracking parameters to your meeting link.
	EngagementsLastMeetingBookedSource string `json:"engagements_last_meeting_booked_source,omitempty"`

	// The date of the most recent meeting (past or upcoming) logged for, scheduled with, or booked by a contact associated with this deal.
	HsLatestMeetingActivity DateTime `json:"hs_latest_meeting_activity,omitempty"`

	// The last time a tracked sales email was replied to for this deal
	HsSalesEmailLastReplied DateTime `json:"hs_sales_email_last_replied,omitempty"`

	// User the deal is assigned to. Assign additional users to a deal record by creating a custom user property.
	HubspotOwnerId Enumeration `json:"hubspot_owner_id,omitempty"`

	// The last time a call, sales email, or meeting was logged for this deal. This is set automatically by HubSpot based on user actions.
	NotesLastContacted DateTime `json:"notes_last_contacted,omitempty"`

	// The last time a note, call, email, meeting, or task was logged for a deal. This is updated automatically by HubSpot.
	NotesLastUpdated DateTime `json:"notes_last_updated,omitempty"`

	// The date of the next upcoming activity for a deal. This property is set automatically by HubSpot based on user action. This includes logging a future call, sales email, or meeting using the Log feature, as well as creating a future task or scheduling a future meeting. This is updated automatically by HubSpot.
	NotesNextActivityDate DateTime `json:"notes_next_activity_date,omitempty"`

	// The total number of sales activities (notes, calls, emails, meetings, or tasks) logged for a deal. This is updated automatically by HubSpot.
	NumContactedNotes Int `json:"num_contacted_notes,omitempty"`

	// The total number of times a sales email or call has been logged for a deal. This is updated automatically by HubSpot.
	NumNotes Int `json:"num_notes,omitempty"`

	// The date the deal was created. This property is set automatically by HubSpot.
	HsCreatedate DateTime `json:"hs_createdate,omitempty"`

	// Primary team of the deal owner. This property is set automatically by HubSpot.
	HubspotTeamId Enumeration `json:"hubspot_team_id,omitempty"`

	// The type of deal. By default, categorize your deal as either a New Business or Existing Business.
	Dealtype Enumeration `json:"dealtype,omitempty"`

	// The value of all owner referencing properties for this object, both default and custom
	HsAllOwnerIds Enumeration `json:"hs_all_owner_ids,omitempty"`

	// Description of the deal
	Description string `json:"description,omitempty"`

	// The team ids corresponding to all owner referencing properties for this object, both default and custom
	HsAllTeamIds Enumeration `json:"hs_all_team_ids,omitempty"`

	// The team ids, including up the team hierarchy, corresponding to all owner referencing properties for this object, both default and custom
	HsAllAccessibleTeamIds Enumeration `json:"hs_all_accessible_team_ids,omitempty"`

	// The number of contacts associated with this deal. This property is set automatically by HubSpot.
	NumAssociatedContacts Int `json:"num_associated_contacts,omitempty"`

	// Reason why this deal was lost
	ClosedLostReason string `json:"closed_lost_reason,omitempty"`

	// Reason why this deal was won
	ClosedWonReason string `json:"closed_won_reason,omitempty"`
}

type FeedbackSubmissionDefaultProperties struct {
	// Conversation agent's email
	HsAgentEmail string `json:"hs_agent_email,omitempty"`

	// Agent responsible for the conversation which led to the feedback
	HsAgentId Int `json:"hs_agent_id,omitempty"`

	// Conversation agent's full name
	HsAgentName string `json:"hs_agent_name,omitempty"`

	// The team IDs, including up the team hierarchy, corresponding to all owner referencing properties for this object, both default and custom
	HsAllAccessibleTeamIds Enumeration `json:"hs_all_accessible_team_ids,omitempty"`

	// The business units this record is assigned to.
	HsAllAssignedBusinessUnitIds Enumeration `json:"hs_all_assigned_business_unit_ids,omitempty"`

	// The value of all owner referencing properties for this object, both default and custom.
	HsAllOwnerIds Enumeration `json:"hs_all_owner_ids,omitempty"`

	// The team ids corresponding to all owner referencing properties for this object, both default and custom.
	HsAllTeamIds Enumeration `json:"hs_all_team_ids,omitempty"`

	// Chatflow ID of the conversation associated to the submission
	HsChatflowId Int `json:"hs_chatflow_id,omitempty"`

	// The name of the chatflow associated to the submission
	HsChatflowName string `json:"hs_chatflow_name,omitempty"`

	// Chatflow object id of the conversation associated to the submission
	HsChatflowObjectId Int `json:"hs_chatflow_object_id,omitempty"`

	// First name of the contact associated with the submission
	HsContactFirstname string `json:"hs_contact_firstname,omitempty"`

	// The id of the contact most recently associated with this submission
	HsContactId Int `json:"hs_contact_id,omitempty"`

	// Last name of the contact associated with the submission
	HsContactLastname string `json:"hs_contact_lastname,omitempty"`

	// Thread id of the conversation associated to the submission
	HsConversationThreadId Int `json:"hs_conversation_thread_id,omitempty"`

	// The user that created this object. This value is automatically set by HubSpot and may not be modified.
	HsCreatedByUserId Int `json:"hs_created_by_user_id,omitempty"`

	// The date and time at which this object was created. This value is automatically set by HubSpot and may not be modified.
	HsCreatedate DateTime `json:"hs_createdate,omitempty"`

	// Most recent timestamp of any property update for this object. This includes HubSpot internal properties, which can be visible or hidden. This property is updated automatically.
	HsLastmodifieddate DateTime `json:"hs_lastmodifieddate,omitempty"`

	// The list of record IDs that have been merged into this record. This value is automatically set by HubSpot and may not be modified.
	HsMergedObjectIds Enumeration `json:"hs_merged_object_ids,omitempty"`

	// The unique ID for this record. This value is automatically set by HubSpot and may not be modified.
	HsObjectId Int `json:"hs_object_id,omitempty"`

	// Source (PropertySource) that created this object record
	HsObjectSource string `json:"hs_object_source,omitempty"`

	// First level of detail on how this record was created
	HsObjectSourceDetail1 string `json:"hs_object_source_detail_1,omitempty"`

	// Second level of detail on how this record was created
	HsObjectSourceDetail2 string `json:"hs_object_source_detail_2,omitempty"`

	// Third level of detail on how this record was created
	HsObjectSourceDetail3 string `json:"hs_object_source_detail_3,omitempty"`

	// The sourceId -- further detail -- of the source that created this object record
	HsObjectSourceId string `json:"hs_object_source_id,omitempty"`

	// How this record was created
	HsObjectSourceLabel Enumeration `json:"hs_object_source_label,omitempty"`

	// User ID of the user who initiated creation of this object record
	HsObjectSourceUserId Int `json:"hs_object_source_user_id,omitempty"`

	// Is the object read only
	HsReadOnly Bool `json:"hs_read_only,omitempty"`

	// The ID of the Survey in CRM
	HsSurveyCrmObjectId Enumeration `json:"hs_survey_crm_object_id,omitempty"`

	// Tag IDs assigned to this submission.
	HsTagIds Enumeration `json:"hs_tag_ids,omitempty"`

	// Tags assigned to this submission.
	HsTags Enumeration `json:"hs_tags,omitempty"`

	// ID of the ticket associated with the submission
	HsTicketId Int `json:"hs_ticket_id,omitempty"`

	// Avatar URI of the ticket associated to the feedback's owner
	HsTicketOwnerAvatarUri string `json:"hs_ticket_owner_avatar_uri,omitempty"`

	// User the ticket is assigned to which is associated to feedback submission
	HsTicketOwnerId Int `json:"hs_ticket_owner_id,omitempty"`

	// Short summary of ticket associated with the submission
	HsTicketSubject string `json:"hs_ticket_subject,omitempty"`

	// Unique property used for idempotent creates
	HsUniqueCreationKey string `json:"hs_unique_creation_key,omitempty"`

	// The user that last updated this object. This value is automatically set by HubSpot and may not be modified.
	HsUpdatedByUserId Int `json:"hs_updated_by_user_id,omitempty"`

	// The user IDs of all users that have clicked follow within the object to opt-in to getting follow notifications
	HsUserIdsOfAllNotificationFollowers Enumeration `json:"hs_user_ids_of_all_notification_followers,omitempty"`

	// The user IDs of all object owners that have clicked unfollow within the object to opt-out of getting follow notifications
	HsUserIdsOfAllNotificationUnfollowers Enumeration `json:"hs_user_ids_of_all_notification_unfollowers,omitempty"`

	// The user IDs of all owners of this object
	HsUserIdsOfAllOwners Enumeration `json:"hs_user_ids_of_all_owners,omitempty"`

	// Object is part of an import
	HsWasImported Bool `json:"hs_was_imported,omitempty"`

	// The most recent date an owner was assigned to this object. This is set automatically by HubSpot and can be used for segmentation and reporting.
	HubspotOwnerAssigneddate DateTime `json:"hubspot_owner_assigneddate,omitempty"`

	// The owner of the object.
	HubspotOwnerId Enumeration `json:"hubspot_owner_id,omitempty"`

	// The primary team of the owner.
	HubspotTeamId Enumeration `json:"hubspot_team_id,omitempty"`

	// Industry standard question type for custom surveys
	HsIndustryStandardQuestionType Enumeration `json:"hs_industry_standard_question_type,omitempty"`

	// The survey id that the feedback submission is linked with
	HsSurveyId Int `json:"hs_survey_id,omitempty"`

	// The type of the survey
	HsSurveyType Enumeration `json:"hs_survey_type,omitempty"`

	// The channel of the survey when the feedback submission occurred
	HsSurveyChannel Enumeration `json:"hs_survey_channel,omitempty"`

	// The timestamp of the feedback submission
	HsSubmissionTimestamp DateTime `json:"hs_submission_timestamp,omitempty"`

	// The value of the feedback submission
	HsValue Int `json:"hs_value,omitempty"`

	// The sentiment of the feedback submission
	HsResponseGroup Enumeration `json:"hs_response_group,omitempty"`

	// The follow-up response of the feedback submission
	HsContent string `json:"hs_content,omitempty"`

	// The identifier we use when we initially capture a feedback submission
	HsIngestionId string `json:"hs_ingestion_id,omitempty"`

	// The id of the knowledge article this submission was for
	HsKnowledgeArticleId Int `json:"hs_knowledge_article_id,omitempty"`

	// The id of the visitor who submitted this feedback
	HsVisitorId Int `json:"hs_visitor_id,omitempty"`

	// The id of the equivalent engagement object
	HsEngagementId Int `json:"hs_engagement_id,omitempty"`

	// The URL of the page the submission was made from
	HsSubmissionUrl string `json:"hs_submission_url,omitempty"`

	// The name of the feedback survey the submission is linked with
	HsSurveyName string `json:"hs_survey_name,omitempty"`

	// The form guid that the feedback submission is linked with
	HsFormGuid string `json:"hs_form_guid,omitempty"`

	// The email of the contact associated with a feedback submission
	HsContactEmailRollup string `json:"hs_contact_email_rollup,omitempty"`

	// The email of the contact associated with a feedback submission
	HsSubmissionName string `json:"hs_submission_name,omitempty"`
}

type LineItemDefaultProperties struct {
	// The amount of a line item
	Amount Int `json:"amount,omitempty"`

	// The date the line item was created
	Createdate DateTime `json:"createdate,omitempty"`

	// Full description of product
	Description string `json:"description,omitempty"`

	// The discount amount applied
	Discount Int `json:"discount,omitempty"`

	// The annual contract value (ACV) of this product
	HsAcv Int `json:"hs_acv,omitempty"`

	// The team IDs, including up the team hierarchy, corresponding to all owner referencing properties for this object, both default and custom
	HsAllAccessibleTeamIds Enumeration `json:"hs_all_accessible_team_ids,omitempty"`

	// The business units this record is assigned to.
	HsAllAssignedBusinessUnitIds Enumeration `json:"hs_all_assigned_business_unit_ids,omitempty"`

	// The value of all owner referencing properties for this object, both default and custom.
	HsAllOwnerIds Enumeration `json:"hs_all_owner_ids,omitempty"`

	// The team ids corresponding to all owner referencing properties for this object, both default and custom.
	HsAllTeamIds Enumeration `json:"hs_all_team_ids,omitempty"`

	// Whether buyer selected quantity is to be enabled or not for a LineItem
	HsAllowBuyerSelectedQuantity Enumeration `json:"hs_allow_buyer_selected_quantity,omitempty"`

	// The annual recurring revenue (ARR) of this product
	HsArr Int `json:"hs_arr,omitempty"`

	// End date of a fixed billing period.
	HsBillingPeriodEndDate Date `json:"hs_billing_period_end_date,omitempty"`

	// Start date of a fixed billing period.
	HsBillingPeriodStartDate Date `json:"hs_billing_period_start_date,omitempty"`

	// Number of days billing should be delayed by. It allows the customers to start billing sometime in the future
	HsBillingStartDelayDays Int `json:"hs_billing_start_delay_days,omitempty"`

	// Number of months billing should be delayed by. It allows the customers to start billing sometime in the future.
	HsBillingStartDelayMonths Int `json:"hs_billing_start_delay_months,omitempty"`

	// The type of billing start delay. It can be a fixed date in the future, relative delay in number of days or months
	HsBillingStartDelayType Enumeration `json:"hs_billing_start_delay_type,omitempty"`

	// The amount that sold goods cost the HubSpot customer
	HsCostOfGoodsSold Int `json:"hs_cost_of_goods_sold,omitempty"`

	// The user that created this object. This value is automatically set by HubSpot and may not be modified.
	HsCreatedByUserId Int `json:"hs_created_by_user_id,omitempty"`

	// The date and time at which this object was created. This value is automatically set by HubSpot and may not be modified.
	HsCreatedate DateTime `json:"hs_createdate,omitempty"`

	// The discount percentage applied
	HsDiscountPercentage Int `json:"hs_discount_percentage,omitempty"`

	// The ID of a line item
	HsExternalId Int `json:"hs_external_id,omitempty"`

	// Product images.
	HsImages string `json:"hs_images,omitempty"`

	// The date any property on this product was modified
	HsLastmodifieddate DateTime `json:"hs_lastmodifieddate,omitempty"`

	// Currency code for the line item.
	HsLineItemCurrencyCode Enumeration `json:"hs_line_item_currency_code,omitempty"`

	// The margin value of this product
	HsMargin Int `json:"hs_margin,omitempty"`

	// The margin annual contract value of this product
	HsMarginAcv Int `json:"hs_margin_acv,omitempty"`

	// The margin annual recurring revenue of this product
	HsMarginArr Int `json:"hs_margin_arr,omitempty"`

	// The margin monthly recurring revenue of this product
	HsMarginMrr Int `json:"hs_margin_mrr,omitempty"`

	// The margin total contract value of this product
	HsMarginTcv Int `json:"hs_margin_tcv,omitempty"`

	// The list of record IDs that have been merged into this record. This value is automatically set by HubSpot and may not be modified.
	HsMergedObjectIds Enumeration `json:"hs_merged_object_ids,omitempty"`

	// The monthly recurring revenue (MRR) of this product
	HsMrr Int `json:"hs_mrr,omitempty"`

	// The unique ID for this record. This value is automatically set by HubSpot and may not be modified.
	HsObjectId Int `json:"hs_object_id,omitempty"`

	// Source (PropertySource) that created this object record
	HsObjectSource string `json:"hs_object_source,omitempty"`

	// First level of detail on how this record was created
	HsObjectSourceDetail1 string `json:"hs_object_source_detail_1,omitempty"`

	// Second level of detail on how this record was created
	HsObjectSourceDetail2 string `json:"hs_object_source_detail_2,omitempty"`

	// Third level of detail on how this record was created
	HsObjectSourceDetail3 string `json:"hs_object_source_detail_3,omitempty"`

	// The sourceId -- further detail -- of the source that created this object record
	HsObjectSourceId string `json:"hs_object_source_id,omitempty"`

	// How this record was created
	HsObjectSourceLabel Enumeration `json:"hs_object_source_label,omitempty"`

	// User ID of the user who initiated creation of this object record
	HsObjectSourceUserId Int `json:"hs_object_source_user_id,omitempty"`

	// The order which the line item appears on the quotes
	HsPositionOnQuote Int `json:"hs_position_on_quote,omitempty"`

	// Amount of line item after applying tax
	HsPostTaxAmount Int `json:"hs_post_tax_amount,omitempty"`

	// The pre-discount amount of a line item
	HsPreDiscountAmount Int `json:"hs_pre_discount_amount,omitempty"`

	// ID of the product this was copied from
	HsProductId Int `json:"hs_product_id,omitempty"`

	// The type of product. By default, categorize your product as either Inventory, Non-Inventory or Service.
	HsProductType Enumeration `json:"hs_product_type,omitempty"`

	// Is the object read only
	HsReadOnly Bool `json:"hs_read_only,omitempty"`

	// Recurring billing end date for a line item
	HsRecurringBillingEndDate Date `json:"hs_recurring_billing_end_date,omitempty"`

	// Number of payments for the given period on a Line Item
	HsRecurringBillingNumberOfPayments Int `json:"hs_recurring_billing_number_of_payments,omitempty"`

	// Product recurring billing duration
	HsRecurringBillingPeriod string `json:"hs_recurring_billing_period,omitempty"`

	// Recurring billing start date for a line item
	HsRecurringBillingStartDate Date `json:"hs_recurring_billing_start_date,omitempty"`

	// If there are fixed number payments or it goes on until cancelled
	HsRecurringBillingTerms Enumeration `json:"hs_recurring_billing_terms,omitempty"`

	// Unique product identifier
	HsSku string `json:"hs_sku,omitempty"`

	// The amount set by Ecommerce sync
	HsSyncAmount Int `json:"hs_sync_amount,omitempty"`

	// Amount of tax calculated based on tax rate for this line item
	HsTaxAmount Int `json:"hs_tax_amount,omitempty"`

	// Represents tax rate as percent out of 100. Supports up to 4 decimal places
	HsTaxRate Int `json:"hs_tax_rate,omitempty"`

	// Describes the type of the tax applied on a LI. For example VAT, Tax, GST
	HsTaxType Enumeration `json:"hs_tax_type,omitempty"`

	// The total contract value (TCV) of this product
	HsTcv Int `json:"hs_tcv,omitempty"`

	// The number of months in the term of this product
	HsTermInMonths Int `json:"hs_term_in_months,omitempty"`

	// Calculated total Discount for the line item taking in consideration the flat amount and discount percentage.
	HsTotalDiscount Int `json:"hs_total_discount,omitempty"`

	// Unique property used for idempotent creates
	HsUniqueCreationKey string `json:"hs_unique_creation_key,omitempty"`

	// The user that last updated this object. This value is automatically set by HubSpot and may not be modified.
	HsUpdatedByUserId Int `json:"hs_updated_by_user_id,omitempty"`

	// Product url.
	HsUrl string `json:"hs_url,omitempty"`

	// The user IDs of all users that have clicked follow within the object to opt-in to getting follow notifications
	HsUserIdsOfAllNotificationFollowers Enumeration `json:"hs_user_ids_of_all_notification_followers,omitempty"`

	// The user IDs of all object owners that have clicked unfollow within the object to opt-out of getting follow notifications
	HsUserIdsOfAllNotificationUnfollowers Enumeration `json:"hs_user_ids_of_all_notification_unfollowers,omitempty"`

	// The user IDs of all owners of this object
	HsUserIdsOfAllOwners Enumeration `json:"hs_user_ids_of_all_owners,omitempty"`

	// Variant id of the shopify product
	HsVariantId Int `json:"hs_variant_id,omitempty"`

	// Object is part of an import
	HsWasImported Bool `json:"hs_was_imported,omitempty"`

	// The most recent date an owner was assigned to this object. This is set automatically by HubSpot and can be used for segmentation and reporting.
	HubspotOwnerAssigneddate DateTime `json:"hubspot_owner_assigneddate,omitempty"`

	// The owner of the object.
	HubspotOwnerId Enumeration `json:"hubspot_owner_id,omitempty"`

	// The primary team of the owner.
	HubspotTeamId Enumeration `json:"hubspot_team_id,omitempty"`

	// Product name
	Name string `json:"name,omitempty"`

	// Cost of product
	Price Int `json:"price,omitempty"`

	// How many units of a product are in this line item
	Quantity Int `json:"quantity,omitempty"`

	// How often a line item with recurring billing is billed. It informs the pricing calculation for deals and quotes. Line items with one-time billing aren’t included.
	Recurringbillingfrequency Enumeration `json:"recurringbillingfrequency,omitempty"`

	// The tax amount applied
	Tax Int `json:"tax,omitempty"`
}

type ProductDefaultProperties struct {
	// Internal placeholder (to prevent conflicts with line item property of same name)
	Amount Int `json:"amount,omitempty"`

	// The date the product was created
	Createdate DateTime `json:"createdate,omitempty"`

	// Full description of product
	Description string `json:"description,omitempty"`

	// The discount amount applied
	Discount Int `json:"discount,omitempty"`

	// The team IDs, including up the team hierarchy, corresponding to all owner referencing properties for this object, both default and custom
	HsAllAccessibleTeamIds Enumeration `json:"hs_all_accessible_team_ids,omitempty"`

	// The business units this record is assigned to.
	HsAllAssignedBusinessUnitIds Enumeration `json:"hs_all_assigned_business_unit_ids,omitempty"`

	// The value of all owner referencing properties for this object, both default and custom.
	HsAllOwnerIds Enumeration `json:"hs_all_owner_ids,omitempty"`

	// The team ids corresponding to all owner referencing properties for this object, both default and custom.
	HsAllTeamIds Enumeration `json:"hs_all_team_ids,omitempty"`

	// The path in the FileManager CDN for this product's image. Automatically set by HubSpot.
	HsAvatarFilemanagerKey string `json:"hs_avatar_filemanager_key,omitempty"`

	// The amount that sold goods cost the HubSpot customer
	HsCostOfGoodsSold Int `json:"hs_cost_of_goods_sold,omitempty"`

	// The user that created this object. This value is automatically set by HubSpot and may not be modified.
	HsCreatedByUserId Int `json:"hs_created_by_user_id,omitempty"`

	// The date and time at which this object was created. This value is automatically set by HubSpot and may not be modified.
	HsCreatedate DateTime `json:"hs_createdate,omitempty"`

	// The discount percentage applied
	HsDiscountPercentage Int `json:"hs_discount_percentage,omitempty"`

	// The folder containing this product
	HsFolder Enumeration `json:"hs_folder,omitempty"`

	// The ID of the folder that has this product
	HsFolderId Int `json:"hs_folder_id,omitempty"`

	// Name of the folder containing this product
	HsFolderName string `json:"hs_folder_name,omitempty"`

	// Product images.
	HsImages string `json:"hs_images,omitempty"`

	// The date any property on this product was modified
	HsLastmodifieddate DateTime `json:"hs_lastmodifieddate,omitempty"`

	// The list of record IDs that have been merged into this record. This value is automatically set by HubSpot and may not be modified.
	HsMergedObjectIds Enumeration `json:"hs_merged_object_ids,omitempty"`

	// The unique ID for this record. This value is automatically set by HubSpot and may not be modified.
	HsObjectId Int `json:"hs_object_id,omitempty"`

	// Source (PropertySource) that created this object record
	HsObjectSource string `json:"hs_object_source,omitempty"`

	// First level of detail on how this record was created
	HsObjectSourceDetail1 string `json:"hs_object_source_detail_1,omitempty"`

	// Second level of detail on how this record was created
	HsObjectSourceDetail2 string `json:"hs_object_source_detail_2,omitempty"`

	// Third level of detail on how this record was created
	HsObjectSourceDetail3 string `json:"hs_object_source_detail_3,omitempty"`

	// The sourceId -- further detail -- of the source that created this object record
	HsObjectSourceId string `json:"hs_object_source_id,omitempty"`

	// How this record was created
	HsObjectSourceLabel Enumeration `json:"hs_object_source_label,omitempty"`

	// User ID of the user who initiated creation of this object record
	HsObjectSourceUserId Int `json:"hs_object_source_user_id,omitempty"`

	// The type of product. By default, categorize your product as either Inventory, Non-Inventory or Service.
	HsProductType Enumeration `json:"hs_product_type,omitempty"`

	// Is the object read only
	HsReadOnly Bool `json:"hs_read_only,omitempty"`

	// Product recurring billing duration
	HsRecurringBillingPeriod string `json:"hs_recurring_billing_period,omitempty"`

	// Recurring billing start date for a line item
	HsRecurringBillingStartDate Date `json:"hs_recurring_billing_start_date,omitempty"`

	// Unique product identifier
	HsSku string `json:"hs_sku,omitempty"`

	// Unique property used for idempotent creates
	HsUniqueCreationKey string `json:"hs_unique_creation_key,omitempty"`

	// The user that last updated this object. This value is automatically set by HubSpot and may not be modified.
	HsUpdatedByUserId Int `json:"hs_updated_by_user_id,omitempty"`

	// Product url.
	HsUrl string `json:"hs_url,omitempty"`

	// The user IDs of all users that have clicked follow within the object to opt-in to getting follow notifications
	HsUserIdsOfAllNotificationFollowers Enumeration `json:"hs_user_ids_of_all_notification_followers,omitempty"`

	// The user IDs of all object owners that have clicked unfollow within the object to opt-out of getting follow notifications
	HsUserIdsOfAllNotificationUnfollowers Enumeration `json:"hs_user_ids_of_all_notification_unfollowers,omitempty"`

	// The user IDs of all owners of this object
	HsUserIdsOfAllOwners Enumeration `json:"hs_user_ids_of_all_owners,omitempty"`

	// Object is part of an import
	HsWasImported Bool `json:"hs_was_imported,omitempty"`

	// The most recent date an owner was assigned to this object. This is set automatically by HubSpot and can be used for segmentation and reporting.
	HubspotOwnerAssigneddate DateTime `json:"hubspot_owner_assigneddate,omitempty"`

	// The owner of the object.
	HubspotOwnerId Enumeration `json:"hubspot_owner_id,omitempty"`

	// The primary team of the owner.
	HubspotTeamId Enumeration `json:"hubspot_team_id,omitempty"`

	// Product name
	Name string `json:"name,omitempty"`

	// Cost of product
	Price Int `json:"price,omitempty"`

	// Internal placeholder (to prevent conflicts with line item property of same name)
	Quantity Int `json:"quantity,omitempty"`

	// How often a product with recurring billing is billed. It informs the pricing calculation for deals and quotes. Products with one-time billing aren’t included.
	Recurringbillingfrequency Enumeration `json:"recurringbillingfrequency,omitempty"`

	// The tax amount applied
	Tax Int `json:"tax,omitempty"`
}

type QuoteDefaultProperties struct {
	// The business units this record is assigned to.
	HsAllAssignedBusinessUnitIds Enumeration `json:"hs_all_assigned_business_unit_ids,omitempty"`

	// Accepted forms of payment if quote is using HubSpot Payment
	HsAllowedPaymentMethods Enumeration `json:"hs_allowed_payment_methods,omitempty"`

	//
	HsApproverId Enumeration `json:"hs_approver_id,omitempty"`

	// Indicates if this quote is archived or not
	HsArchived Bool `json:"hs_archived,omitempty"`

	// Indicates if billing address should be collected with this quote
	HsCollectBillingAddress Bool `json:"hs_collect_billing_address,omitempty"`

	// The user that created this object. This value is automatically set by HubSpot and may not be modified.
	HsCreatedByUserId Int `json:"hs_created_by_user_id,omitempty"`

	// The date the quote was created
	HsCreatedate DateTime `json:"hs_createdate,omitempty"`

	// Indicates if total contract value (TCV) should be displayed on the published quote
	HsDisplayTcvOnQuote Bool `json:"hs_display_tcv_on_quote,omitempty"`

	// Domain this quote should be served from
	HsDomain string `json:"hs_domain,omitempty"`

	// The date and time the document was signed by all signers
	HsEsignDate DateTime `json:"hs_esign_date,omitempty"`

	// Path of the expiration template for the quote
	HsExpirationTemplatePath string `json:"hs_expiration_template_path,omitempty"`

	// Feedback about this quote during the approval process
	HsFeedback string `json:"hs_feedback,omitempty"`

	// The language for displaying this quote publicly
	HsLanguage string `json:"hs_language,omitempty"`

	// The last time any property on the quote was modified
	HsLastmodifieddate DateTime `json:"hs_lastmodifieddate,omitempty"`

	// For internal HubSpot Application use only. Global term for the discount percentage applied.
	HsLineItemGlobalTermHsDiscountPercentage string `json:"hs_line_item_global_term_hs_discount_percentage,omitempty"`

	// For internal HubSpot Application use only. Indicates if the Global term for the discount percentage is enabled.
	HsLineItemGlobalTermHsDiscountPercentageEnabled Bool `json:"hs_line_item_global_term_hs_discount_percentage_enabled,omitempty"`

	// For internal HubSpot Application use only. Global term for product recurring billing duration.
	HsLineItemGlobalTermHsRecurringBillingPeriod string `json:"hs_line_item_global_term_hs_recurring_billing_period,omitempty"`

	// For internal HubSpot Application use only. Indicates if the Global term for product recurring billing duration is enabled.
	HsLineItemGlobalTermHsRecurringBillingPeriodEnabled Bool `json:"hs_line_item_global_term_hs_recurring_billing_period_enabled,omitempty"`

	// For internal HubSpot Application use only. Global term for recurring billing start date for a line item.
	HsLineItemGlobalTermHsRecurringBillingStartDate string `json:"hs_line_item_global_term_hs_recurring_billing_start_date,omitempty"`

	// For internal HubSpot Application use only. Indicates if the Global term for recurring billing start date for a line item is enabled.
	HsLineItemGlobalTermHsRecurringBillingStartDateEnabled Bool `json:"hs_line_item_global_term_hs_recurring_billing_start_date_enabled,omitempty"`

	// For internal HubSpot Application use only. Global term for how frequently the product is billed.
	HsLineItemGlobalTermRecurringbillingfrequency string `json:"hs_line_item_global_term_recurringbillingfrequency,omitempty"`

	// For internal HubSpot Application use only. Indicates if the Global term for how frequently the product is billed is enabled
	HsLineItemGlobalTermRecurringbillingfrequencyEnabled Bool `json:"hs_line_item_global_term_recurringbillingfrequency_enabled,omitempty"`

	// The locale for displaying this quote publicly
	HsLocale string `json:"hs_locale,omitempty"`

	// Indicates if this quote is locked and can never be updated
	HsLocked Bool `json:"hs_locked,omitempty"`

	// Indicates if the written signature quote is manually signed
	HsManuallySigned Bool `json:"hs_manually_signed,omitempty"`

	// The list of record IDs that have been merged into this record. This value is automatically set by HubSpot and may not be modified.
	HsMergedObjectIds Enumeration `json:"hs_merged_object_ids,omitempty"`

	// The unique ID for this record. This value is automatically set by HubSpot and may not be modified.
	HsObjectId Int `json:"hs_object_id,omitempty"`

	// Source (PropertySource) that created this object record
	HsObjectSource string `json:"hs_object_source,omitempty"`

	// First level of detail on how this record was created
	HsObjectSourceDetail1 string `json:"hs_object_source_detail_1,omitempty"`

	// Second level of detail on how this record was created
	HsObjectSourceDetail2 string `json:"hs_object_source_detail_2,omitempty"`

	// Third level of detail on how this record was created
	HsObjectSourceDetail3 string `json:"hs_object_source_detail_3,omitempty"`

	// The sourceId -- further detail -- of the source that created this object record
	HsObjectSourceId string `json:"hs_object_source_id,omitempty"`

	// How this record was created
	HsObjectSourceLabel Enumeration `json:"hs_object_source_label,omitempty"`

	// User ID of the user who initiated creation of this object record
	HsObjectSourceUserId Int `json:"hs_object_source_user_id,omitempty"`

	// The date the quote was paid by the customer
	HsPaymentDate DateTime `json:"hs_payment_date,omitempty"`

	// Payment status for the quote
	HsPaymentStatus Enumeration `json:"hs_payment_status,omitempty"`

	// Payment type for the quote if payment is enabled
	HsPaymentType Enumeration `json:"hs_payment_type,omitempty"`

	// The link to download a pdf of a quote
	HsPdfDownloadLink string `json:"hs_pdf_download_link,omitempty"`

	//
	HsPdfGenerationStatus Enumeration `json:"hs_pdf_generation_status,omitempty"`

	// Domain this proposal should be served from
	HsProposalDomain string `json:"hs_proposal_domain,omitempty"`

	// Path to serve this proposal from
	HsProposalSlug string `json:"hs_proposal_slug,omitempty"`

	// Template Path this proposal is rendered with
	HsProposalTemplatePath string `json:"hs_proposal_template_path,omitempty"`

	// Key for accessing quote document URL
	HsPublicUrlKey string `json:"hs_public_url_key,omitempty"`

	// Represents the url on which a quote is hosted
	HsQuoteLink string `json:"hs_quote_link,omitempty"`

	// What LineItem property will be used to calculate the Quote Amount
	HsQuoteTotalPreference Enumeration `json:"hs_quote_total_preference,omitempty"`

	// Is the object read only
	HsReadOnly Bool `json:"hs_read_only,omitempty"`

	// The render status of a quote
	HsRenderStatus Enumeration `json:"hs_render_status,omitempty"`

	// URL to image of the sender avatar image to be set on the quote
	HsSenderImageUrl string `json:"hs_sender_image_url,omitempty"`

	//
	HsSignStatus Enumeration `json:"hs_sign_status,omitempty"`

	// Path to serve this quote from
	HsSlug string `json:"hs_slug,omitempty"`

	// The total contract value (TCV) of this quote.
	HsTcv Int `json:"hs_tcv,omitempty"`

	// Whether this Quote is rendered with the default templates
	HsTemplateType Enumeration `json:"hs_template_type,omitempty"`

	// Indicates if the quote is in test mode
	HsTestMode Bool `json:"hs_test_mode,omitempty"`

	// Unique property used for idempotent creates
	HsUniqueCreationKey string `json:"hs_unique_creation_key,omitempty"`

	// The user that last updated this object. This value is automatically set by HubSpot and may not be modified.
	HsUpdatedByUserId Int `json:"hs_updated_by_user_id,omitempty"`

	// The user IDs of all users that have clicked follow within the object to opt-in to getting follow notifications
	HsUserIdsOfAllNotificationFollowers Enumeration `json:"hs_user_ids_of_all_notification_followers,omitempty"`

	// The user IDs of all object owners that have clicked unfollow within the object to opt-out of getting follow notifications
	HsUserIdsOfAllNotificationUnfollowers Enumeration `json:"hs_user_ids_of_all_notification_unfollowers,omitempty"`

	// The user IDs of all owners of this object
	HsUserIdsOfAllOwners Enumeration `json:"hs_user_ids_of_all_owners,omitempty"`

	// Object is part of an import
	HsWasImported Bool `json:"hs_was_imported,omitempty"`

	// The date an owner was assigned to the quote
	HubspotOwnerAssigneddate DateTime `json:"hubspot_owner_assigneddate,omitempty"`

	// The name of the company sending this quote
	HsSenderCompanyName string `json:"hs_sender_company_name,omitempty"`

	// First name of the sender of this quote
	HsSenderFirstname string `json:"hs_sender_firstname,omitempty"`

	// The title of this quote
	HsTitle string `json:"hs_title,omitempty"`

	// The date that this quote expires
	HsExpirationDate DateTime `json:"hs_expiration_date,omitempty"`

	// The domain of the company sending this quote
	HsSenderCompanyDomain string `json:"hs_sender_company_domain,omitempty"`

	// Last name of the sender of this quote
	HsSenderLastname string `json:"hs_sender_lastname,omitempty"`

	// Comments to the buyer
	HsComments string `json:"hs_comments,omitempty"`

	// The street address of the company sending this quote
	HsSenderCompanyAddress string `json:"hs_sender_company_address,omitempty"`

	// Email address of the sender of this quote
	HsSenderEmail string `json:"hs_sender_email,omitempty"`

	// The second line of the street address of the company sending this quote
	HsSenderCompanyAddress2 string `json:"hs_sender_company_address2,omitempty"`

	// Phone number of the sender of this quote
	HsSenderPhone string `json:"hs_sender_phone,omitempty"`

	// Any relevant information about pricing, purchasing terms, and/or master agreements
	HsTerms string `json:"hs_terms,omitempty"`

	// URL to image of the logo displayed on the quote
	HsLogoUrl string `json:"hs_logo_url,omitempty"`

	// The city in which the company sending this quote is located
	HsSenderCompanyCity string `json:"hs_sender_company_city,omitempty"`

	// Job title of the sender of this quote
	HsSenderJobtitle string `json:"hs_sender_jobtitle,omitempty"`

	// The state/region in which the company sending this quote is located
	HsSenderCompanyState string `json:"hs_sender_company_state,omitempty"`

	// Indicates if the quote document should have a place for a written signature
	HsShowSignatureBox Bool `json:"hs_show_signature_box,omitempty"`

	// Color value used on the quote document
	HsPrimaryColor string `json:"hs_primary_color,omitempty"`

	// The last time a tracked sales email was replied to for this quote
	HsSalesEmailLastReplied DateTime `json:"hs_sales_email_last_replied,omitempty"`

	// The zip code in which the company sending this quote is located
	HsSenderCompanyZip string `json:"hs_sender_company_zip,omitempty"`

	// Indicates if the quote document should have a place for a written countersignature
	HsShowCountersignatureBox Bool `json:"hs_show_countersignature_box,omitempty"`

	// The owner of the quote
	HubspotOwnerId Enumeration `json:"hubspot_owner_id,omitempty"`

	// The last timestamp when a call, email or meeting was logged for this quote
	NotesLastContacted DateTime `json:"notes_last_contacted,omitempty"`

	// The last time a note, call, email, meeting, or task was logged for a quote. This is set automatically by HubSpot based on user actions in the quote record.
	NotesLastUpdated DateTime `json:"notes_last_updated,omitempty"`

	// The date of the next upcoming activity for this quote
	NotesNextActivityDate DateTime `json:"notes_next_activity_date,omitempty"`

	// The number of times a call, email or meeting was logged for this quote
	NumContactedNotes Int `json:"num_contacted_notes,omitempty"`

	// Number of sales activities for this quote
	NumNotes Int `json:"num_notes,omitempty"`

	// Currency code for the quote
	HsCurrency string `json:"hs_currency,omitempty"`

	// The country in which the company sending this quote is located
	HsSenderCompanyCountry string `json:"hs_sender_company_country,omitempty"`

	// The name of the team associated with the owner of the quote
	HubspotTeamId Enumeration `json:"hubspot_team_id,omitempty"`

	// The value of all owner referencing properties for this object, both default and custom
	HsAllOwnerIds Enumeration `json:"hs_all_owner_ids,omitempty"`

	// URL to image of the sender company image to be set on the quote
	HsSenderCompanyImageUrl string `json:"hs_sender_company_image_url,omitempty"`

	// Timezone for displaying dates on the quote
	HsTimezone string `json:"hs_timezone,omitempty"`

	// The team ids corresponding to all owner referencing properties for this object, both default and custom
	HsAllTeamIds Enumeration `json:"hs_all_team_ids,omitempty"`

	// Indicates if payment can be collected via the quote link
	HsPaymentEnabled Bool `json:"hs_payment_enabled,omitempty"`

	// The team ids, including up the team hierarchy, corresponding to all owner referencing properties for this object, both default and custom
	HsAllAccessibleTeamIds Enumeration `json:"hs_all_accessible_team_ids,omitempty"`

	// Indicates if esign is enabled for quote
	HsEsignEnabled Bool `json:"hs_esign_enabled,omitempty"`

	// The number of deals associated with this quote. This is set automatically by HubSpot.
	HsNumAssociatedDeals Int `json:"hs_num_associated_deals,omitempty"`

	// Total number of signers required for this quote to be fully signed
	HsEsignNumSignersRequired Int `json:"hs_esign_num_signers_required,omitempty"`

	// The legacy quote template used to generate a quote. These templates will be sunset on Dec. 12th 2022.
	HsTemplate Enumeration `json:"hs_template,omitempty"`

	// Number of completed signatures collected so far
	HsEsignNumSignersCompleted Int `json:"hs_esign_num_signers_completed,omitempty"`

	// The total due for the quote
	HsQuoteAmount Int `json:"hs_quote_amount,omitempty"`

	// Approval status of the quote
	HsStatus Enumeration `json:"hs_status,omitempty"`

	// Reference number shown on quote document
	HsQuoteNumber string `json:"hs_quote_number,omitempty"`

	// Indicates if shipping address should be collected with this quote
	HsCollectShippingAddress Bool `json:"hs_collect_shipping_address,omitempty"`
}

type DiscountDefaultProperties struct {
	// The team IDs, including up the team hierarchy, corresponding to all owner referencing properties for this object, both default and custom
	HsAllAccessibleTeamIds Enumeration `json:"hs_all_accessible_team_ids,omitempty"`

	// The business units this record is assigned to.
	HsAllAssignedBusinessUnitIds Enumeration `json:"hs_all_assigned_business_unit_ids,omitempty"`

	// The value of all owner referencing properties for this object, both default and custom.
	HsAllOwnerIds Enumeration `json:"hs_all_owner_ids,omitempty"`

	// The team ids corresponding to all owner referencing properties for this object, both default and custom.
	HsAllTeamIds Enumeration `json:"hs_all_team_ids,omitempty"`

	// The user that created this object. This value is automatically set by HubSpot and may not be modified.
	HsCreatedByUserId Int `json:"hs_created_by_user_id,omitempty"`

	// The date and time at which this object was created. This value is automatically set by HubSpot and may not be modified.
	HsCreatedate DateTime `json:"hs_createdate,omitempty"`

	// How often the discount will be applied.
	HsDuration Enumeration `json:"hs_duration,omitempty"`

	// The name of the discount.
	HsLabel string `json:"hs_label,omitempty"`

	// Most recent timestamp of any property update for this object. This includes HubSpot internal properties, which can be visible or hidden. This property is updated automatically.
	HsLastmodifieddate DateTime `json:"hs_lastmodifieddate,omitempty"`

	// The list of record IDs that have been merged into this record. This value is automatically set by HubSpot and may not be modified.
	HsMergedObjectIds Enumeration `json:"hs_merged_object_ids,omitempty"`

	// The unique ID for this record. This value is automatically set by HubSpot and may not be modified.
	HsObjectId Int `json:"hs_object_id,omitempty"`

	// Source (PropertySource) that created this object record
	HsObjectSource string `json:"hs_object_source,omitempty"`

	// First level of detail on how this record was created
	HsObjectSourceDetail1 string `json:"hs_object_source_detail_1,omitempty"`

	// Second level of detail on how this record was created
	HsObjectSourceDetail2 string `json:"hs_object_source_detail_2,omitempty"`

	// Third level of detail on how this record was created
	HsObjectSourceDetail3 string `json:"hs_object_source_detail_3,omitempty"`

	// The sourceId -- further detail -- of the source that created this object record
	HsObjectSourceId string `json:"hs_object_source_id,omitempty"`

	// How this record was created
	HsObjectSourceLabel Enumeration `json:"hs_object_source_label,omitempty"`

	// User ID of the user who initiated creation of this object record
	HsObjectSourceUserId Int `json:"hs_object_source_user_id,omitempty"`

	// Is the object read only
	HsReadOnly Bool `json:"hs_read_only,omitempty"`

	// The order in which the discount will be displayed and applied. Changing this order affects how the total discount is calculated.
	HsSortOrder Int `json:"hs_sort_order,omitempty"`

	// How the discount will be calculated, and what it will apply to. It could be a fixed amount of the total, a percentage of the total, or a fixed amount per unit.
	HsType Enumeration `json:"hs_type,omitempty"`

	// Unique property used for idempotent creates
	HsUniqueCreationKey string `json:"hs_unique_creation_key,omitempty"`

	// The user that last updated this object. This value is automatically set by HubSpot and may not be modified.
	HsUpdatedByUserId Int `json:"hs_updated_by_user_id,omitempty"`

	// The user IDs of all users that have clicked follow within the object to opt-in to getting follow notifications
	HsUserIdsOfAllNotificationFollowers Enumeration `json:"hs_user_ids_of_all_notification_followers,omitempty"`

	// The user IDs of all object owners that have clicked unfollow within the object to opt-out of getting follow notifications
	HsUserIdsOfAllNotificationUnfollowers Enumeration `json:"hs_user_ids_of_all_notification_unfollowers,omitempty"`

	// The user IDs of all owners of this object
	HsUserIdsOfAllOwners Enumeration `json:"hs_user_ids_of_all_owners,omitempty"`

	// The monetary or percentage value of the discount.
	HsValue Int `json:"hs_value,omitempty"`

	// Object is part of an import
	HsWasImported Bool `json:"hs_was_imported,omitempty"`

	// The most recent date an owner was assigned to this object. This is set automatically by HubSpot and can be used for segmentation and reporting.
	HubspotOwnerAssigneddate DateTime `json:"hubspot_owner_assigneddate,omitempty"`

	// The owner of the object.
	HubspotOwnerId Enumeration `json:"hubspot_owner_id,omitempty"`

	// The primary team of the owner.
	HubspotTeamId Enumeration `json:"hubspot_team_id,omitempty"`
}

type FeeDefaultProperties struct {
	// The team IDs, including up the team hierarchy, corresponding to all owner referencing properties for this object, both default and custom
	HsAllAccessibleTeamIds Enumeration `json:"hs_all_accessible_team_ids,omitempty"`

	// The business units this record is assigned to.
	HsAllAssignedBusinessUnitIds Enumeration `json:"hs_all_assigned_business_unit_ids,omitempty"`

	// The value of all owner referencing properties for this object, both default and custom.
	HsAllOwnerIds Enumeration `json:"hs_all_owner_ids,omitempty"`

	// The team ids corresponding to all owner referencing properties for this object, both default and custom.
	HsAllTeamIds Enumeration `json:"hs_all_team_ids,omitempty"`

	// The user that created this object. This value is automatically set by HubSpot and may not be modified.
	HsCreatedByUserId Int `json:"hs_created_by_user_id,omitempty"`

	// The date and time at which this object was created. This value is automatically set by HubSpot and may not be modified.
	HsCreatedate DateTime `json:"hs_createdate,omitempty"`

	// The name of the fee.
	HsLabel string `json:"hs_label,omitempty"`

	// Most recent timestamp of any property update for this object. This includes HubSpot internal properties, which can be visible or hidden. This property is updated automatically.
	HsLastmodifieddate DateTime `json:"hs_lastmodifieddate,omitempty"`

	// The list of record IDs that have been merged into this record. This value is automatically set by HubSpot and may not be modified.
	HsMergedObjectIds Enumeration `json:"hs_merged_object_ids,omitempty"`

	// The unique ID for this record. This value is automatically set by HubSpot and may not be modified.
	HsObjectId Int `json:"hs_object_id,omitempty"`

	// Source (PropertySource) that created this object record
	HsObjectSource string `json:"hs_object_source,omitempty"`

	// First level of detail on how this record was created
	HsObjectSourceDetail1 string `json:"hs_object_source_detail_1,omitempty"`

	// Second level of detail on how this record was created
	HsObjectSourceDetail2 string `json:"hs_object_source_detail_2,omitempty"`

	// Third level of detail on how this record was created
	HsObjectSourceDetail3 string `json:"hs_object_source_detail_3,omitempty"`

	// The sourceId -- further detail -- of the source that created this object record
	HsObjectSourceId string `json:"hs_object_source_id,omitempty"`

	// How this record was created
	HsObjectSourceLabel Enumeration `json:"hs_object_source_label,omitempty"`

	// User ID of the user who initiated creation of this object record
	HsObjectSourceUserId Int `json:"hs_object_source_user_id,omitempty"`

	// Is the object read only
	HsReadOnly Bool `json:"hs_read_only,omitempty"`

	// The order in which the fee will be displayed.
	HsSortOrder Int `json:"hs_sort_order,omitempty"`

	// How the fee will be calculated: as a fixed amount or a percentage.
	HsType Enumeration `json:"hs_type,omitempty"`

	// Unique property used for idempotent creates
	HsUniqueCreationKey string `json:"hs_unique_creation_key,omitempty"`

	// The user that last updated this object. This value is automatically set by HubSpot and may not be modified.
	HsUpdatedByUserId Int `json:"hs_updated_by_user_id,omitempty"`

	// The user IDs of all users that have clicked follow within the object to opt-in to getting follow notifications
	HsUserIdsOfAllNotificationFollowers Enumeration `json:"hs_user_ids_of_all_notification_followers,omitempty"`

	// The user IDs of all object owners that have clicked unfollow within the object to opt-out of getting follow notifications
	HsUserIdsOfAllNotificationUnfollowers Enumeration `json:"hs_user_ids_of_all_notification_unfollowers,omitempty"`

	// The user IDs of all owners of this object
	HsUserIdsOfAllOwners Enumeration `json:"hs_user_ids_of_all_owners,omitempty"`

	// The monetary or percentage value of the fee.
	HsValue Int `json:"hs_value,omitempty"`

	// Object is part of an import
	HsWasImported Bool `json:"hs_was_imported,omitempty"`

	// The most recent date an owner was assigned to this object. This is set automatically by HubSpot and can be used for segmentation and reporting.
	HubspotOwnerAssigneddate DateTime `json:"hubspot_owner_assigneddate,omitempty"`

	// The owner of the object.
	HubspotOwnerId Enumeration `json:"hubspot_owner_id,omitempty"`

	// The primary team of the owner.
	HubspotTeamId Enumeration `json:"hubspot_team_id,omitempty"`
}

type TaxDefaultProperties struct {
	// The team IDs, including up the team hierarchy, corresponding to all owner referencing properties for this object, both default and custom
	HsAllAccessibleTeamIds Enumeration `json:"hs_all_accessible_team_ids,omitempty"`

	// The business units this record is assigned to.
	HsAllAssignedBusinessUnitIds Enumeration `json:"hs_all_assigned_business_unit_ids,omitempty"`

	// The value of all owner referencing properties for this object, both default and custom.
	HsAllOwnerIds Enumeration `json:"hs_all_owner_ids,omitempty"`

	// The team ids corresponding to all owner referencing properties for this object, both default and custom.
	HsAllTeamIds Enumeration `json:"hs_all_team_ids,omitempty"`

	// The user that created this object. This value is automatically set by HubSpot and may not be modified.
	HsCreatedByUserId Int `json:"hs_created_by_user_id,omitempty"`

	// The date and time at which this object was created. This value is automatically set by HubSpot and may not be modified.
	HsCreatedate DateTime `json:"hs_createdate,omitempty"`

	// The name of the tax.
	HsLabel string `json:"hs_label,omitempty"`

	// Most recent timestamp of any property update for this object. This includes HubSpot internal properties, which can be visible or hidden. This property is updated automatically.
	HsLastmodifieddate DateTime `json:"hs_lastmodifieddate,omitempty"`

	// The list of record IDs that have been merged into this record. This value is automatically set by HubSpot and may not be modified.
	HsMergedObjectIds Enumeration `json:"hs_merged_object_ids,omitempty"`

	// The unique ID for this record. This value is automatically set by HubSpot and may not be modified.
	HsObjectId Int `json:"hs_object_id,omitempty"`

	// Source (PropertySource) that created this object record
	HsObjectSource string `json:"hs_object_source,omitempty"`

	// First level of detail on how this record was created
	HsObjectSourceDetail1 string `json:"hs_object_source_detail_1,omitempty"`

	// Second level of detail on how this record was created
	HsObjectSourceDetail2 string `json:"hs_object_source_detail_2,omitempty"`

	// Third level of detail on how this record was created
	HsObjectSourceDetail3 string `json:"hs_object_source_detail_3,omitempty"`

	// The sourceId -- further detail -- of the source that created this object record
	HsObjectSourceId string `json:"hs_object_source_id,omitempty"`

	// How this record was created
	HsObjectSourceLabel Enumeration `json:"hs_object_source_label,omitempty"`

	// User ID of the user who initiated creation of this object record
	HsObjectSourceUserId Int `json:"hs_object_source_user_id,omitempty"`

	// Is the object read only
	HsReadOnly Bool `json:"hs_read_only,omitempty"`

	// The order in which the tax will be displayed.
	HsSortOrder Int `json:"hs_sort_order,omitempty"`

	// How the tax will be calculated: as a fixed amount or a percentage.
	HsType Enumeration `json:"hs_type,omitempty"`

	// Unique property used for idempotent creates
	HsUniqueCreationKey string `json:"hs_unique_creation_key,omitempty"`

	// The user that last updated this object. This value is automatically set by HubSpot and may not be modified.
	HsUpdatedByUserId Int `json:"hs_updated_by_user_id,omitempty"`

	// The user IDs of all users that have clicked follow within the object to opt-in to getting follow notifications
	HsUserIdsOfAllNotificationFollowers Enumeration `json:"hs_user_ids_of_all_notification_followers,omitempty"`

	// The user IDs of all object owners that have clicked unfollow within the object to opt-out of getting follow notifications
	HsUserIdsOfAllNotificationUnfollowers Enumeration `json:"hs_user_ids_of_all_notification_unfollowers,omitempty"`

	// The user IDs of all owners of this object
	HsUserIdsOfAllOwners Enumeration `json:"hs_user_ids_of_all_owners,omitempty"`

	// The monetary or percentage value of the tax.
	HsValue Int `json:"hs_value,omitempty"`

	// Object is part of an import
	HsWasImported Bool `json:"hs_was_imported,omitempty"`

	// The most recent date an owner was assigned to this object. This is set automatically by HubSpot and can be used for segmentation and reporting.
	HubspotOwnerAssigneddate DateTime `json:"hubspot_owner_assigneddate,omitempty"`

	// The owner of the object.
	HubspotOwnerId Enumeration `json:"hubspot_owner_id,omitempty"`

	// The primary team of the owner.
	HubspotTeamId Enumeration `json:"hubspot_team_id,omitempty"`
}

type TicketDefaultProperties struct {
	// The date the ticket was closed
	ClosedDate DateTime `json:"closed_date,omitempty"`

	// VID of contact that created the ticket
	CreatedBy Int `json:"created_by,omitempty"`

	// The date the ticket was created
	Createdate DateTime `json:"createdate,omitempty"`

	// The date of the first email response from an agent after a ticket was created
	FirstAgentReplyDate DateTime `json:"first_agent_reply_date,omitempty"`

	// The business units this record is assigned to.
	HsAllAssignedBusinessUnitIds Enumeration `json:"hs_all_assigned_business_unit_ids,omitempty"`

	// All associated contact companies
	HsAllAssociatedContactCompanies string `json:"hs_all_associated_contact_companies,omitempty"`

	// All associated contact emails
	HsAllAssociatedContactEmails string `json:"hs_all_associated_contact_emails,omitempty"`

	// All associated contact first names
	HsAllAssociatedContactFirstnames string `json:"hs_all_associated_contact_firstnames,omitempty"`

	// All associated contact last names
	HsAllAssociatedContactLastnames string `json:"hs_all_associated_contact_lastnames,omitempty"`

	// All associated contact mobile phones
	HsAllAssociatedContactMobilephones string `json:"hs_all_associated_contact_mobilephones,omitempty"`

	// All associated contact phones
	HsAllAssociatedContactPhones string `json:"hs_all_associated_contact_phones,omitempty"`

	// All mentioned users on the associated conversations
	HsAllConversationMentions Enumeration `json:"hs_all_conversation_mentions,omitempty"`

	// Reference to the SLA rule that was applied to this ticket
	HsAppliedSlaRuleConfigId Int `json:"hs_applied_sla_rule_config_id,omitempty"`

	// teamIds of the ticket owner(s)’ main team(s), or if ticket owner is blank, then teamId(s) of the team(s) we attempted to assign to
	HsAssignedTeamIds Enumeration `json:"hs_assigned_team_ids,omitempty"`

	// defines how the Object Assignment is done
	HsAssignmentMethod Enumeration `json:"hs_assignment_method,omitempty"`

	// Thread that this ticket was automatically created for using ticket rules
	HsAutoGeneratedFromThreadId Int `json:"hs_auto_generated_from_thread_id,omitempty"`

	// Conversations Message Id of the message that originated this ticket
	HsConversationsOriginatingMessageId string `json:"hs_conversations_originating_message_id,omitempty"`

	// Thread that this ticket was originally created for
	HsConversationsOriginatingThreadId Int `json:"hs_conversations_originating_thread_id,omitempty"`

	// The user that created this object. This value is automatically set by HubSpot and may not be modified.
	HsCreatedByUserId Int `json:"hs_created_by_user_id,omitempty"`

	// Internal read-only property representing the date the ticket was created in HubSpot
	HsCreatedate DateTime `json:"hs_createdate,omitempty"`

	// ID of the custom inbox the ticket is associated with
	HsCustomInbox Int `json:"hs_custom_inbox,omitempty"`

	// The date and time when the ticket entered the 'New' stage, 'Support Pipeline' pipeline
	HsDateEntered1 DateTime `json:"hs_date_entered_1,omitempty"`

	// The date and time when the ticket entered the 'Waiting on contact' stage, 'Support Pipeline' pipeline
	HsDateEntered2 DateTime `json:"hs_date_entered_2,omitempty"`

	// The date and time when the ticket entered the 'Waiting on us' stage, 'Support Pipeline' pipeline
	HsDateEntered3 DateTime `json:"hs_date_entered_3,omitempty"`

	// The date and time when the ticket entered the 'Closed' stage, 'Support Pipeline' pipeline
	HsDateEntered4 DateTime `json:"hs_date_entered_4,omitempty"`

	// The date and time when the ticket exited the 'New' stage, 'Support Pipeline' pipeline
	HsDateExited1 DateTime `json:"hs_date_exited_1,omitempty"`

	// The date and time when the ticket exited the 'Waiting on contact' stage, 'Support Pipeline' pipeline
	HsDateExited2 DateTime `json:"hs_date_exited_2,omitempty"`

	// The date and time when the ticket exited the 'Waiting on us' stage, 'Support Pipeline' pipeline
	HsDateExited3 DateTime `json:"hs_date_exited_3,omitempty"`

	// The date and time when the ticket exited the 'Closed' stage, 'Support Pipeline' pipeline
	HsDateExited4 DateTime `json:"hs_date_exited_4,omitempty"`

	// Unique ids corresponding to tickets in a system outside of HubSpot
	HsExternalObjectIds Enumeration `json:"hs_external_object_ids,omitempty"`

	// Last CES survey comment that this contact gave for this ticket
	HsFeedbackLastCesFollowUp string `json:"hs_feedback_last_ces_follow_up,omitempty"`

	// Last CES survey rating that this contact gave for this ticket
	HsFeedbackLastCesRating Enumeration `json:"hs_feedback_last_ces_rating,omitempty"`

	// The time that this contact last submitted a CES survey response. This is automatically set by HubSpot.
	HsFeedbackLastSurveyDate DateTime `json:"hs_feedback_last_survey_date,omitempty"`

	// Files attached to a support form by a contact.
	HsFileUpload string `json:"hs_file_upload,omitempty"`

	// The date of the first response from an agent out of all associated conversations
	HsFirstAgentMessageSentAt DateTime `json:"hs_first_agent_message_sent_at,omitempty"`

	// A calculated property to help with sorting in the Helpdesk
	HsHelpdeskSortTimestamp DateTime `json:"hs_helpdesk_sort_timestamp,omitempty"`

	// Is this Ticket rendered in the Help Desk
	HsInHelpdesk Bool `json:"hs_in_helpdesk,omitempty"`

	// Inbox the ticket is in
	HsInboxId Int `json:"hs_inbox_id,omitempty"`

	// Whether the ticket is visible in help desk
	HsIsVisibleInHelpDesk Bool `json:"hs_is_visible_in_help_desk,omitempty"`

	// The type of the last email activity with the contact associated with the ticket.
	HsLastEmailActivity Enumeration `json:"hs_last_email_activity,omitempty"`

	// The date of the last email activity with the contact associated with the ticket.
	HsLastEmailDate DateTime `json:"hs_last_email_date,omitempty"`

	// Whether the last message came from visitor
	HsLastMessageFromVisitor Bool `json:"hs_last_message_from_visitor,omitempty"`

	// The date of the last response from the visitor
	HsLastMessageReceivedAt DateTime `json:"hs_last_message_received_at,omitempty"`

	// The date of the last response from an agent or bot
	HsLastMessageSentAt DateTime `json:"hs_last_message_sent_at,omitempty"`

	// The last time a note, call, email, meeting, or task was logged for a ticket. This is updated automatically by HubSpot.
	HsLastactivitydate DateTime `json:"hs_lastactivitydate,omitempty"`

	// The last time a call, chat conversation, LinkedIn message, postal mail, meeting, sales email, SMS, or WhatsApp message was logged for a ticket. This is set automatically by HubSpot based on user actions in the ticket record.
	HsLastcontacted DateTime `json:"hs_lastcontacted,omitempty"`

	// Most recent timestamp of any property update for this ticket. This includes HubSpot internal properties, which can be visible or hidden. This property is updated automatically.
	HsLastmodifieddate DateTime `json:"hs_lastmodifieddate,omitempty"`

	// Agents who have seen the newest message across all conversations associated to the ticket
	HsLatestMessageSeenByAgentIds Enumeration `json:"hs_latest_message_seen_by_agent_ids,omitempty"`

	// The list of Ticket record IDs that have been merged into this Ticket. This value is automatically set by HubSpot and may not be modified.
	HsMergedObjectIds Enumeration `json:"hs_merged_object_ids,omitempty"`

	// Most relevant sla status
	HsMostRelevantSlaStatus Enumeration `json:"hs_most_relevant_sla_status,omitempty"`

	// Most relevant SLA type between Close By, First Response, Next Response
	HsMostRelevantSlaType Enumeration `json:"hs_most_relevant_sla_type,omitempty"`

	// Microsoft Teams message ID for this ticket.
	HsMsteamsMessageId string `json:"hs_msteams_message_id,omitempty"`

	// The date of the next upcoming activity for a ticket. This property is set automatically by HubSpot based on user action. This includes logging a future call, email, or meeting using the Log feature, as well as creating a future task or scheduling a future meeting. This is updated automatically by HubSpot.
	HsNextactivitydate DateTime `json:"hs_nextactivitydate,omitempty"`

	// Number of companies associated with this ticket
	HsNumAssociatedCompanies Int `json:"hs_num_associated_companies,omitempty"`

	// Number of conversations associated to the ticket
	HsNumAssociatedConversations Int `json:"hs_num_associated_conversations,omitempty"`

	// The number of times a call, email, or meeting was logged on the ticket
	HsNumTimesContacted Int `json:"hs_num_times_contacted,omitempty"`

	// The unique ID for this record. This value is automatically set by HubSpot and may not be modified.
	HsObjectId Int `json:"hs_object_id,omitempty"`

	// Source (PropertySource) that created this object record
	HsObjectSource string `json:"hs_object_source,omitempty"`

	// First level of detail on how this record was created
	HsObjectSourceDetail1 string `json:"hs_object_source_detail_1,omitempty"`

	// Second level of detail on how this record was created
	HsObjectSourceDetail2 string `json:"hs_object_source_detail_2,omitempty"`

	// Third level of detail on how this record was created
	HsObjectSourceDetail3 string `json:"hs_object_source_detail_3,omitempty"`

	// The sourceId -- further detail -- of the source that created this object record
	HsObjectSourceId string `json:"hs_object_source_id,omitempty"`

	// How this record was created
	HsObjectSourceLabel Enumeration `json:"hs_object_source_label,omitempty"`

	// User ID of the user who initiated creation of this object record
	HsObjectSourceUserId Int `json:"hs_object_source_user_id,omitempty"`

	// First channel account used when conversation was started
	HsOriginatingChannelInstanceId Enumeration `json:"hs_originating_channel_instance_id,omitempty"`

	// Engagement id of the email originating this ticket
	HsOriginatingEmailEngagementId Int `json:"hs_originating_email_engagement_id,omitempty"`

	// The channel the conversation is in
	HsOriginatingGenericChannelId Enumeration `json:"hs_originating_generic_channel_id,omitempty"`

	// The object ID of the current pinned engagement. This value is automatically set by HubSpot and may not be modified.
	HsPinnedEngagementId Int `json:"hs_pinned_engagement_id,omitempty"`

	// The pipeline that contains this ticket
	HsPipeline Enumeration `json:"hs_pipeline,omitempty"`

	// The pipeline stage that contains this ticket
	HsPipelineStage Enumeration `json:"hs_pipeline_stage,omitempty"`

	// Primary company of a ticket
	HsPrimaryCompany string `json:"hs_primary_company,omitempty"`

	// Primary company ID of a ticket
	HsPrimaryCompanyId Int `json:"hs_primary_company_id,omitempty"`

	// Primary company name of a ticket
	HsPrimaryCompanyName string `json:"hs_primary_company_name,omitempty"`

	// Is the object read only
	HsReadOnly Bool `json:"hs_read_only,omitempty"`

	// The action taken to resolve the ticket
	HsResolution Enumeration `json:"hs_resolution,omitempty"`

	// Stores a list of users who have viewed the most recent interaction on a ticket
	HsSeenByAgentIds Enumeration `json:"hs_seen_by_agent_ids,omitempty"`

	// The ID of the object from which the data was migrated. This is set automatically during portal data migration.
	HsSourceObjectId Int `json:"hs_source_object_id,omitempty"`

	// List of tag ids applicable to a ticket. This property is set automatically by HubSpot.
	HsTagIds Enumeration `json:"hs_tag_ids,omitempty"`

	// Thread IDs (from cv-threads) used to implement custom cascading delete/restore
	HsThreadIdsToRestore Enumeration `json:"hs_thread_ids_to_restore,omitempty"`

	// Main reason customer reached out for help
	HsTicketCategory Enumeration `json:"hs_ticket_category,omitempty"`

	// The unique id for this ticket. This unique id is automatically populated by HubSpot.
	HsTicketId Int `json:"hs_ticket_id,omitempty"`

	// The level of attention needed on the ticket
	HsTicketPriority Enumeration `json:"hs_ticket_priority,omitempty"`

	// The total time in seconds spent by the ticket in the 'New' stage, 'Support Pipeline' pipeline
	HsTimeIn1 Int `json:"hs_time_in_1,omitempty"`

	// The total time in seconds spent by the ticket in the 'Waiting on contact' stage, 'Support Pipeline' pipeline
	HsTimeIn2 Int `json:"hs_time_in_2,omitempty"`

	// The total time in seconds spent by the ticket in the 'Waiting on us' stage, 'Support Pipeline' pipeline
	HsTimeIn3 Int `json:"hs_time_in_3,omitempty"`

	// The total time in seconds spent by the ticket in the 'Closed' stage, 'Support Pipeline' pipeline
	HsTimeIn4 Int `json:"hs_time_in_4,omitempty"`

	// When the ticket falls out of Time to Close SLA.
	HsTimeToCloseSlaAt DateTime `json:"hs_time_to_close_sla_at,omitempty"`

	// Current Time to Close SLA status of ticket
	HsTimeToCloseSlaStatus Enumeration `json:"hs_time_to_close_sla_status,omitempty"`

	// When the ticket falls out of the Time to First Response SLA
	HsTimeToFirstResponseSlaAt DateTime `json:"hs_time_to_first_response_sla_at,omitempty"`

	// Current Time to First Response SLA status.
	HsTimeToFirstResponseSlaStatus Enumeration `json:"hs_time_to_first_response_sla_status,omitempty"`

	// When the ticket falls out of Time to Next Response SLA
	HsTimeToNextResponseSlaAt DateTime `json:"hs_time_to_next_response_sla_at,omitempty"`

	// Current Time to Next Response SLA status
	HsTimeToNextResponseSlaStatus Enumeration `json:"hs_time_to_next_response_sla_status,omitempty"`

	// Unique property used for idempotent creates
	HsUniqueCreationKey string `json:"hs_unique_creation_key,omitempty"`

	// The user that last updated this object. This value is automatically set by HubSpot and may not be modified.
	HsUpdatedByUserId Int `json:"hs_updated_by_user_id,omitempty"`

	// The user IDs of all users that have clicked follow within the object to opt-in to getting follow notifications
	HsUserIdsOfAllNotificationFollowers Enumeration `json:"hs_user_ids_of_all_notification_followers,omitempty"`

	// The user IDs of all object owners that have clicked unfollow within the object to opt-out of getting follow notifications
	HsUserIdsOfAllNotificationUnfollowers Enumeration `json:"hs_user_ids_of_all_notification_unfollowers,omitempty"`

	// The user IDs of all owners of this object
	HsUserIdsOfAllOwners Enumeration `json:"hs_user_ids_of_all_owners,omitempty"`

	// Object is part of an import
	HsWasImported Bool `json:"hs_was_imported,omitempty"`

	// The date an owner was assigned to the ticket
	HubspotOwnerAssigneddate DateTime `json:"hubspot_owner_assigneddate,omitempty"`

	// The date of the last reply or note
	LastEngagementDate DateTime `json:"last_engagement_date,omitempty"`

	// The date of the last customer response
	LastReplyDate DateTime `json:"last_reply_date,omitempty"`

	// Answer to NPS follow up question
	NpsFollowUpAnswer string `json:"nps_follow_up_answer,omitempty"`

	// Specific version of NPS follow up question that was asked
	NpsFollowUpQuestionVersion Int `json:"nps_follow_up_question_version,omitempty"`

	// NPS score received after ticket resolution
	NpsScore Enumeration `json:"nps_score,omitempty"`

	// The id of an email thread with ticket conversation
	SourceThreadId string `json:"source_thread_id,omitempty"`

	// The time between when the ticket was created and closed
	TimeToClose Int `json:"time_to_close,omitempty"`

	// The time from the ticket create date to the first agent email reply
	TimeToFirstAgentReply Int `json:"time_to_first_agent_reply,omitempty"`

	// Short summary of ticket
	Subject string `json:"subject,omitempty"`

	// Description of the ticket
	Content string `json:"content,omitempty"`

	// Channel where ticket was originally submitted
	SourceType Enumeration `json:"source_type,omitempty"`

	// The id of a connected source object
	SourceRef string `json:"source_ref,omitempty"`

	// Tags associated with your tickets
	Tags Enumeration `json:"tags,omitempty"`

	// The last time a tracked sales email was replied to for this ticket
	HsSalesEmailLastReplied DateTime `json:"hs_sales_email_last_replied,omitempty"`

	// User the ticket is assigned to. Assign additional users to a ticket record by creating a custom user property.
	HubspotOwnerId Enumeration `json:"hubspot_owner_id,omitempty"`

	// The last time a call, chat conversation, LinkedIn message, postal mail, meeting, note, sales email, SMS, or WhatsApp message was logged for a ticket. This is set automatically by HubSpot based on user actions in the ticket record.
	NotesLastContacted DateTime `json:"notes_last_contacted,omitempty"`

	// The last time a note, call, email, meeting, or task was logged for a ticket. This is set automatically by HubSpot based on user actions in the ticket record.
	NotesLastUpdated DateTime `json:"notes_last_updated,omitempty"`

	// The date of the next upcoming activity for this ticket
	NotesNextActivityDate DateTime `json:"notes_next_activity_date,omitempty"`

	// The number of times a call, email or meeting was logged for this ticket
	NumContactedNotes Int `json:"num_contacted_notes,omitempty"`

	// The number of times a call, chat conversation, LinkedIn message, postal mail, meeting, note, sales email, SMS, task, or WhatsApp message was logged for a ticket record. This is set automatically by HubSpot based on user actions in the ticket record.
	NumNotes Int `json:"num_notes,omitempty"`

	// Primary team of the ticket owner. This property is set automatically by HubSpot.
	HubspotTeamId Enumeration `json:"hubspot_team_id,omitempty"`

	// The value of all owner referencing properties for this object, both default and custom
	HsAllOwnerIds Enumeration `json:"hs_all_owner_ids,omitempty"`

	// The team IDs, including up the team hierarchy, corresponding to all owner referencing properties for this object, both default and custom
	HsAllTeamIds Enumeration `json:"hs_all_team_ids,omitempty"`

	// The team IDs, including up the team hierarchy, corresponding to all owner referencing properties for this object, both default and custom
	HsAllAccessibleTeamIds Enumeration `json:"hs_all_accessible_team_ids,omitempty"`
}

type GoalDefaultProperties struct {
	// Describes if the goal target can be treated as deleted.
	Hs_MigrationSoftDelete Enumeration `json:"hs__migration_soft_delete,omitempty"`

	// A list of ';'-separated universally unique IDs of ad accounts, each formatted: {AD_NETWORK}-{ACCOUNT_ID}
	HsAdAccountAssetIds Enumeration `json:"hs_ad_account_asset_ids,omitempty"`

	// A list of ';'-separated universally unique IDs of ad campaigns, each formatted: {NETWORK}-{CAMPAIGN_ID}
	HsAdCampaignAssetIds Enumeration `json:"hs_ad_campaign_asset_ids,omitempty"`

	// The team IDs, including up the team hierarchy, corresponding to all owner referencing properties for this object, both default and custom
	HsAllAccessibleTeamIds Enumeration `json:"hs_all_accessible_team_ids,omitempty"`

	// The business units this record is assigned to.
	HsAllAssignedBusinessUnitIds Enumeration `json:"hs_all_assigned_business_unit_ids,omitempty"`

	// The value of all owner referencing properties for this object, both default and custom.
	HsAllOwnerIds Enumeration `json:"hs_all_owner_ids,omitempty"`

	// Which team owns the goal (there should only ever be one team)
	HsAllTeamIds Enumeration `json:"hs_all_team_ids,omitempty"`

	// The team ID of the assigned team of the goal (unpopulated if the goal is not team-assigned or team-by-user-assigned)
	HsAssigneeTeamId Int `json:"hs_assignee_team_id,omitempty"`

	// Goal assignee type calculated field based on the assigned user or team ids on the goal
	HsAssigneeType Enumeration `json:"hs_assignee_type,omitempty"`

	// The user ID of the assigned user of the goal (unpopulated if the goal is not user-assigned)
	HsAssigneeUserId Int `json:"hs_assignee_user_id,omitempty"`

	// List of marketing campaign asset ids
	HsCampaignAssetIds Enumeration `json:"hs_campaign_asset_ids,omitempty"`

	// The lifecycle stage selected for the contact lifecycle goal
	HsContactLifecycleStage string `json:"hs_contact_lifecycle_stage,omitempty"`

	// The user that created this object. This value is automatically set by HubSpot and may not be modified.
	HsCreatedByUserId Int `json:"hs_created_by_user_id,omitempty"`

	// The date and time at which this object was created. This value is automatically set by HubSpot and may not be modified.
	HsCreatedate DateTime `json:"hs_createdate,omitempty"`

	// DEPRECATED - DO NOT USE! Code for the currency conversion
	HsCurrency string `json:"hs_currency,omitempty"`

	// Exchange rate to be used for calculated properties reflecting currency-based values in home currencies. Value is automatically set for in-progress goals. Once a goal has ended, the value should be alterable by admins only.
	HsCurrencyExchangeRate Int `json:"hs_currency_exchange_rate,omitempty"`

	// A serialized list of DEAL pipeline IDs used to filter the KPI value calculation.
	HsDealPipelineIds Enumeration `json:"hs_deal_pipeline_ids,omitempty"`

	// The frequency of goal edit update notifications
	HsEditUpdatesNotificationFrequency Enumeration `json:"hs_edit_updates_notification_frequency,omitempty"`

	// The last day covered by the Goal Target's date range.
	HsEndDate Date `json:"hs_end_date,omitempty"`

	// The last day covered by the Goal Target's date range.
	HsEndDatetime DateTime `json:"hs_end_datetime,omitempty"`

	// Represents the number of months the fiscal year is away from a standard calendar year (0 means the fiscal year starts in January).
	HsFiscalYearOffset Int `json:"hs_fiscal_year_offset,omitempty"`

	// Calculated key based on several goal kpi properties which define a goal. Example goal definition key: 0-3-amount_in_home_currency-SUM-hs_closed_date-[]-HIGHER_IS_BETTER-goalName1
	HsGoalDefinitionKey string `json:"hs_goal_definition_key,omitempty"`

	// Calculated key based on several goal kpi properties which define a goal including the team id of a goal. Example goal definition key: 0-3-amount_in_home_currency-SUM-[filterGroup with hubspot_team_id]-HIGHER_IS_BETTER-goalName1
	HsGoalDefinitionKeyWithTeam string `json:"hs_goal_definition_key_with_team,omitempty"`

	// Calculated key based on several goal kpi properties which define a goal including the user id of a goal. Example goal definition key: 0-3-amount_in_home_currency-SUM-[filterGroup]-HIGHER_IS_BETTER-goalName1-userId1
	HsGoalDefinitionKeyWithUser string `json:"hs_goal_definition_key_with_user,omitempty"`

	// The name of the goal.
	HsGoalName string `json:"hs_goal_name,omitempty"`

	// Currency code for the goal target
	HsGoalTargetCurrencyCode Enumeration `json:"hs_goal_target_currency_code,omitempty"`

	// Object ID of the associated Goal Target Group
	HsGoalTargetGroupId Int `json:"hs_goal_target_group_id,omitempty"`

	// The goal type
	HsGoalType Enumeration `json:"hs_goal_type,omitempty"`

	// Unique ID set on all Goal Targets and Goal Target Groups created by a single request
	HsGroupCorrelationUuid string `json:"hs_group_correlation_uuid,omitempty"`

	// READ ME! (kind of) indicates whether the goal should appear in the forecasting app - this property may be true even if the goal should NOT appear in the forecasting app. As of April 2023, this property will be true if hs_milestone is either  \“monthly\” or \“quarterly\“, and if the count of pipelines in hs_pipelines 0 or 1. No other checks are made!
	HsIsForecastable Enumeration `json:"hs_is_forecastable,omitempty"`

	// Describes if the object was created during the migration process and it comes from the old app.
	HsIsLegacy Enumeration `json:"hs_is_legacy,omitempty"`

	// Display units for goals based on duration properties
	HsKpiDisplayUnit Enumeration `json:"hs_kpi_display_unit,omitempty"`

	// Filter groups for KPI calculations
	HsKpiFilterGroups string `json:"hs_kpi_filter_groups,omitempty"`

	// Filter groups which excludes pipelines for use with goal definition key
	HsKpiFilterGroupsForKeyGrouping string `json:"hs_kpi_filter_groups_for_key_grouping,omitempty"`

	// Filter groups with the exclusion of pipeline and team filters for use with team goal definition keys
	HsKpiFilterGroupsForKeyTeamGrouping string `json:"hs_kpi_filter_groups_for_key_team_grouping,omitempty"`

	// Indicates whether this is a rollup goal: whether the KPI allows contributions from child teams (in addition to the assigned team)
	HsKpiIsTeamRollup Bool `json:"hs_kpi_is_team_rollup,omitempty"`

	// Type of metric used to aggregate the property name
	HsKpiMetricType string `json:"hs_kpi_metric_type,omitempty"`

	// The name of the CrmObject the KPI values are based on
	HsKpiObjectType string `json:"hs_kpi_object_type,omitempty"`

	// The ObjectTypeId of the object the KPI values are based on
	HsKpiObjectTypeId string `json:"hs_kpi_object_type_id,omitempty"`

	// Name of the property on the object type aggregated as the KPI value.
	HsKpiPropertyName string `json:"hs_kpi_property_name,omitempty"`

	// Single object custom goal components (metric type, object type, property name)
	HsKpiSingleObjectCustomGoalTypeName string `json:"hs_kpi_single_object_custom_goal_type_name,omitempty"`

	// Time period property
	HsKpiTimePeriodProperty string `json:"hs_kpi_time_period_property,omitempty"`

	// Tracking method for progress directionality
	HsKpiTrackingMethod string `json:"hs_kpi_tracking_method,omitempty"`

	// Unit type of the kpi calculation
	HsKpiUnitType Enumeration `json:"hs_kpi_unit_type,omitempty"`

	// The latest KPI value for this goal target
	HsKpiValue Int `json:"hs_kpi_value,omitempty"`

	// The time that latest KPI value was calculated
	HsKpiValueCalculatedAt Int `json:"hs_kpi_value_calculated_at,omitempty"`

	// The time that latest KPI value was calculated
	HsKpiValueLastCalculatedAt DateTime `json:"hs_kpi_value_last_calculated_at,omitempty"`

	// Most recent timestamp of any property update for this object. This includes HubSpot internal properties, which can be visible or hidden. This property is updated automatically.
	HsLastmodifieddate DateTime `json:"hs_lastmodifieddate,omitempty"`

	// Migrated value describing if the goal is active or soft deleted.
	HsLegacyActive Enumeration `json:"hs_legacy_active,omitempty"`

	// Migrated value of created at.
	HsLegacyCreatedAt Int `json:"hs_legacy_created_at,omitempty"`

	// Migrated value of created by.
	HsLegacyCreatedBy Int `json:"hs_legacy_created_by,omitempty"`

	// Concatenated value of the 3 sql_ids that form a quarterly GT separated by a '-'. Used to be able to update both monthly and quarterly legacy Goal Targets.
	HsLegacyQuarterlyTargetCompositeId string `json:"hs_legacy_quarterly_target_composite_id,omitempty"`

	// Id of the migrated goal target in the old system. This has been deprecated for legacy_unique_sql_id.
	HsLegacySqlId Int `json:"hs_legacy_sql_id,omitempty"`

	// Id of the migrated goal target in the old system. This property has a unique value to ensure no duplicate migrated goals.
	HsLegacyUniqueSqlId Int `json:"hs_legacy_unique_sql_id,omitempty"`

	// Migrated value of updated at.
	HsLegacyUpdatedAt Int `json:"hs_legacy_updated_at,omitempty"`

	// Migrated value of updated by.
	HsLegacyUpdatedBy Int `json:"hs_legacy_updated_by,omitempty"`

	// The list of record IDs that have been merged into this record. This value is automatically set by HubSpot and may not be modified.
	HsMergedObjectIds Enumeration `json:"hs_merged_object_ids,omitempty"`

	// Describes if the goal target can be treated as deleted.
	HsMigrationSoftDelete Enumeration `json:"hs_migration_soft_delete,omitempty"`

	// The period during which a Goal Target is active
	HsMilestone Enumeration `json:"hs_milestone,omitempty"`

	// The unique ID for this record. This value is automatically set by HubSpot and may not be modified.
	HsObjectId Int `json:"hs_object_id,omitempty"`

	// Source (PropertySource) that created this object record
	HsObjectSource string `json:"hs_object_source,omitempty"`

	// First level of detail on how this record was created
	HsObjectSourceDetail1 string `json:"hs_object_source_detail_1,omitempty"`

	// Second level of detail on how this record was created
	HsObjectSourceDetail2 string `json:"hs_object_source_detail_2,omitempty"`

	// Third level of detail on how this record was created
	HsObjectSourceDetail3 string `json:"hs_object_source_detail_3,omitempty"`

	// The sourceId -- further detail -- of the source that created this object record
	HsObjectSourceId string `json:"hs_object_source_id,omitempty"`

	// How this record was created
	HsObjectSourceLabel Enumeration `json:"hs_object_source_label,omitempty"`

	// User ID of the user who initiated creation of this object record
	HsObjectSourceUserId Int `json:"hs_object_source_user_id,omitempty"`

	// Goal Target outcome
	HsOutcome Enumeration `json:"hs_outcome,omitempty"`

	// The owner IDs of all owners of this object
	HsOwnerIdsOfAllOwners Enumeration `json:"hs_owner_ids_of_all_owners,omitempty"`

	// List of deal pipelines that the quota applies to. Used to calculate kpi values
	HsPipelineIds Enumeration `json:"hs_pipeline_ids,omitempty"`

	// A serialized list of pipeline IDs used to filter the KPI value calculation
	HsPipelines string `json:"hs_pipelines,omitempty"`

	// The frequency of goal progress update notifications
	HsProgressUpdatesNotificationFrequency Enumeration `json:"hs_progress_updates_notification_frequency,omitempty"`

	// Is the object read only
	HsReadOnly Bool `json:"hs_read_only,omitempty"`

	// Opt into notifications when goal is achieved
	HsShouldNotifyOnAchieved Enumeration `json:"hs_should_notify_on_achieved,omitempty"`

	// Opt into goal edit update notifications
	HsShouldNotifyOnEditUpdates Enumeration `json:"hs_should_notify_on_edit_updates,omitempty"`

	// Opt into notifications when goal is exceeded
	HsShouldNotifyOnExceeded Enumeration `json:"hs_should_notify_on_exceeded,omitempty"`

	// Opt into notifications when a goal is kicked off
	HsShouldNotifyOnKickoff Enumeration `json:"hs_should_notify_on_kickoff,omitempty"`

	// Opt into notifications when a goal is missed
	HsShouldNotifyOnMissed Enumeration `json:"hs_should_notify_on_missed,omitempty"`

	// Opt into goal progress updates notifications
	HsShouldNotifyOnProgressUpdates Enumeration `json:"hs_should_notify_on_progress_updates,omitempty"`

	// Property used to flag Goal targets that need to be recalculated
	HsShouldRecalculate Enumeration `json:"hs_should_recalculate,omitempty"`

	// The first day that goal target's date range covers
	HsStartDatetime DateTime `json:"hs_start_datetime,omitempty"`

	// Goal Target Filter Group(s) for all filters except time-based filters
	HsStaticKpiFilterGroups string `json:"hs_static_kpi_filter_groups,omitempty"`

	// The status of the Goal Target
	HsStatus Enumeration `json:"hs_status,omitempty"`

	// The goal status display order.
	HsStatusDisplayOrder Int `json:"hs_status_display_order,omitempty"`

	// The target amount for this goal target
	HsTargetAmount Int `json:"hs_target_amount,omitempty"`

	// The amount of the goal target, using the exchange rate, in your company's currency
	HsTargetAmountInHomeCurrency Int `json:"hs_target_amount_in_home_currency,omitempty"`

	// The team ID associated with the Goal Target
	HsTeamId Int `json:"hs_team_id,omitempty"`

	// ID of the template used to create the goal target, if applicable.  HS-defined templates use GoalType enum ordinals for the ID (see HsTemplateFactory).
	HsTemplateId Int `json:"hs_template_id,omitempty"`

	// A serialized list of TICKET pipeline IDs used to filter the KPI value calculation.
	HsTicketPipelineIds Enumeration `json:"hs_ticket_pipeline_ids,omitempty"`

	// Unique property used for idempotent creates
	HsUniqueCreationKey string `json:"hs_unique_creation_key,omitempty"`

	// The user that last updated this object. This value is automatically set by HubSpot and may not be modified.
	HsUpdatedByUserId Int `json:"hs_updated_by_user_id,omitempty"`

	// The user ID associated with the Goal Target
	HsUserId Int `json:"hs_user_id,omitempty"`

	// The user IDs of all users that have clicked follow within the object to opt-in to getting follow notifications
	HsUserIdsOfAllNotificationFollowers Enumeration `json:"hs_user_ids_of_all_notification_followers,omitempty"`

	// The user IDs of all object owners that have clicked unfollow within the object to opt-out of getting follow notifications
	HsUserIdsOfAllNotificationUnfollowers Enumeration `json:"hs_user_ids_of_all_notification_unfollowers,omitempty"`

	// The user IDs of all owners of this object
	HsUserIdsOfAllOwners Enumeration `json:"hs_user_ids_of_all_owners,omitempty"`

	// Object is part of an import
	HsWasImported Bool `json:"hs_was_imported,omitempty"`

	// The most recent date an owner was assigned to this object. This is set automatically by HubSpot and can be used for segmentation and reporting.
	HubspotOwnerAssigneddate DateTime `json:"hubspot_owner_assigneddate,omitempty"`

	// The owner of the object.
	HubspotOwnerId Enumeration `json:"hubspot_owner_id,omitempty"`

	// The primary team of the owner.
	HubspotTeamId Enumeration `json:"hubspot_team_id,omitempty"`

	// The first day that goal target's date range covers
	HsStartDate Date `json:"hs_start_date,omitempty"`

	// Progress percent is represented by the kpi value as a percent of the target amount
	HsKpiProgressPercent Int `json:"hs_kpi_progress_percent,omitempty"`
}
