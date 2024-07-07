# YouTerm Notes

## YouTube API

### Resources

#### Channel

##### GET /channels (list) - Returns a collections of 0+ channel resources

###### Required Request Parameters

part
- auditDetails
- brandingSettings
- contentDetails
- contentOwnerDetails
- id
- localizations
- snippet
- statistics
- status
- topicDetails

###### Filters (Exactly 1 of the following)

- forHandle: Channel's handle (ex: GoogleForDevelopers)
- forUsername: Channel's YouTube username
- id: Channel's YouTube ID (comma delimited list to request multiple)
- mine: set to `true` to get all channels owned by the authenticated user

###### Optional Parameters

- maxResults: 0-50, inclusive. Default is 5
- pageToken: Specific page in a result set to be loaded. prev/next page tokens are returned in paginated searches

###### Errors

- 400 Invalid Criteria
- 403 Channel Forbidden (likely improperly authorized)
- 404 Channel Not Found

##### Resource Schema

- auditDetails
    - Overall Good Standing (boolean)
    - Community Guidelines Good Standing (boolean)
    - Copyright Strikes Good Standing (boolean)
    - Content ID Claim Good Standing (boolean)
- brandingSettings
    - channel
        - title
        - description
        - keywords
        - trackingAnalyticsAccountId
        - unsubscribedTrailer
        - defaultLanguage
        - country
- contentDetails
    - relatedPlaylists
        - likes
        - uploads
- contentOwnerDetails
    - contentOwner
    - timeLinked
- id
- localizations
- snippet
    - title
    - description
    - customUrl
    - publishedAt
    - thumbnails
    - defaultLanguage
    - country
- statistics
    - viewCount (ulong)
    - subscriberCount (ulong)
    - hiddenSubscriberCount (boolean)
    - videoCount (ulong)
- status
    - privacyStatus
    - isLinked
    - longUploadsStatus
    - madeForKids
    - selfDeclaredMadeForKids
- topicDetails
    - topicCategories[] (list of wikipedia URLs about content)

#### Playlist

##### Resource Schema

- id
- snippet
    - publishedAt
    - channelId
    - title
    - description
    - thumbnails
    - channelTitle
- status
    - privacyStatus
- contentDetails
    - itemCount
- player
    - embedHtml

#### PlaylistItem

##### Resource Schema

- id
- snippet
    - publishedAt
    - channelId
    - title
    - description
    - channelTitle
    - playlistId
    - position
- contentDetails
    - videoId
    - startAt
    - endAt
    - note
    - videoPublishedAt
- status
    - privacyStatus


#### Search

#### Subscription

#### Video

##### Resource Schema

- id
- snippet
    - publishedAt
    - channelId
    - title
    - description
    - thumbnails
    - channelTitle
    - tags[]
    - categoryId
    - liveBroadcastContent
    - defaultLanguage
- contentDetails
    - duration
    - dimension
    - definition
    - caption
    - licensedContent
    - regionRestriction
    - contentRating
    - projection
    - hasCustomThumbnail
- status
    - uploadStatus
    - failureReason
    - rejectionReason
    - privacyStatus
    - publishAt
    - license
    - embeddable
    - publicStatsViewable
    - madeForKids
    - selfDeclaredMadeForKids
- statistics
    - viewCount
    - likeCount
    - dislikeCount
    - commentCount
- topicDetails
    - topicCategories[]
- fileDetails
    - fileName
    - fileSize
    - fileType
    - container
    - creationTime

### Resources Not Documented Here

- Activities
- Captions
- ChannelBanners
- ChannelSections
- Comments
- CommentThreads
- Members
- MembershipLevels
- VideoAbuseReportReasons
- Watermarks

### Example Requests/Responses

#### Get details about Channel by handle "Northernlion"

Request: `https://www.googleapis.com/youtube/v3/channels?part=contentDetails&forHandle=Northernlion&key=<API>`

Response:
```json
{
  "kind": "youtube#channelListResponse",
  "etag": "Bm7tWdMGf1SFZIn46TrFiMy7uj8",
  "pageInfo": {
    "totalResults": 1,
    "resultsPerPage": 5
  },
  "items": [
    {
      "kind": "youtube#channel",
      "etag": "5E_MZbXV6zl9Sf4US-ENcHQ36vg",
      "id": "UC3tNpTOHsTnkmbwztCs30sA",
      "contentDetails": {
        "relatedPlaylists": {
          "likes": "",
          "uploads": "UU3tNpTOHsTnkmbwztCs30sA"
        }
      }
    }
  ]
}
```

#### Get details about Playlist by ID

Request: `https://www.googleapis.com/youtube/v3/playlists?part=contentDetails&id=UU3tNpTOHsTnkmbwztCs30sA&key=<KEY>`

Response:
```json
{
    "kind": "youtube#playlistListResponse",
    "etag": "EbPCG5NCUEx_5yYut2YG_epyVhs",
    "pageInfo": {
        "totalResults": 1,
        "resultsPerPage": 5
    },
    "items": [
        {
            "kind": "youtube#playlist",
            "etag": "XMNd_3RV34xFoRx6x3BM9RKh3PE",
            "id": "UU3tNpTOHsTnkmbwztCs30sA",
            "contentDetails": {
                "itemCount": 19923
            }
        }
    ]
}
```

#### Get a list of PlaylistItem resources based on Playlist ID

Request: `https://www.googleapis.com/youtube/v3/playlistItems?part=contentDetails&playlistId=UU3tNpTOHsTnkmbwztCs30sA&key=<KEY>`

Response:
```json
{
    "kind": "youtube#playlistItemListResponse",
    "etag": "6EOaluLC_EEJgwbUAdrCXnPHI_Q",
    "nextPageToken": "EAAaelBUOkNBVWlFRE5FUkRoR09VSkZNMFUyUkVVMFJUVW9BVWpRMEsyWXhPT0dBMUFCV2pZaVEyaG9WbFpVVGpCVWJrSlZWREJvZWxaSE5YSmlWMG96Wlc1U1JHTjZUWGRqTUVWVFEzZHBTM2xOUzNwQ2FFTkJjV1UwTkNJ",
    "items": [
        {
            "kind": "youtube#playlistItem",
            "etag": "S4IgbZcnOA93odpfvuZjz9PK7dk",
            "id": "VVUzdE5wVE9Ic1Rua21id3p0Q3MzMHNBLlpfbEMzLUtBem5r",
            "contentDetails": {
                "videoId": "Z_lC3-KAznk",
                "videoPublishedAt": "2024-06-17T21:00:18Z"
            }
        },
        {
            "kind": "youtube#playlistItem",
            "etag": "UD3O1LCMyl13kdokdIhBHh4eD-E",
            "id": "VVUzdE5wVE9Ic1Rua21id3p0Q3MzMHNBLktyYkd5VVdYUy00",
            "contentDetails": {
                "videoId": "KrbGyUWXS-4",
                "videoPublishedAt": "2024-06-17T20:00:07Z"
            }
        },
        {
            "kind": "youtube#playlistItem",
            "etag": "vgOat0KgIXNk8AxGsmDvJWnkRfA",
            "id": "VVUzdE5wVE9Ic1Rua21id3p0Q3MzMHNBLjViTXBRUFBqdzJv",
            "contentDetails": {
                "videoId": "5bMpQPPjw2o",
                "videoPublishedAt": "2024-06-17T20:00:02Z"
            }
        },
        {
            "kind": "youtube#playlistItem",
            "etag": "2FaAZvyP-QtvXfzqLw7uQmQwnHw",
            "id": "VVUzdE5wVE9Ic1Rua21id3p0Q3MzMHNBLll1WDlQTXpIVnpF",
            "contentDetails": {
                "videoId": "YuX9PMzHVzE",
                "videoPublishedAt": "2024-06-16T20:00:09Z"
            }
        },
        {
            "kind": "youtube#playlistItem",
            "etag": "uESsV9vRAiwjrIxtj1VsZbZgAt8",
            "id": "VVUzdE5wVE9Ic1Rua21id3p0Q3MzMHNBLlVRdU1BV2ZaVVhV",
            "contentDetails": {
                "videoId": "UQuMAWfZUXU",
                "videoPublishedAt": "2024-06-16T16:34:37Z"
            }
        }
    ],
    "pageInfo": {
        "totalResults": 19923,
        "resultsPerPage": 5
    }
}
```

### Code Snippets

#### Get All Playlist Items
```go
package main

import (
        "fmt"
        "log"

        "google.golang.org/api/youtube/v3"
)

func playlistItemsList(service *youtube.Service, part string, playlistId string, pageToken string) *youtube.PlaylistItemListResponse {
        call := service.PlaylistItems.List(part)
        call = call.PlaylistId(playlistId)
        if pageToken != "" {
                call = call.PageToken(pageToken)
        }
        response, err := call.Do()
        handleError(err, "")
        return response
}
```
