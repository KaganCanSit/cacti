/**
 *
 * Please note:
 * This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * Do not edit this file manually.
 *
 */

@file:Suppress(
    "ArrayInDataClass",
    "EnumEntryName",
    "RemoveRedundantQualifierName",
    "UnusedImport"
)

package org.openapitools.client.models

import org.openapitools.client.models.CactiBlockTransactionEventV1

import com.squareup.moshi.Json
import com.squareup.moshi.JsonClass

/**
 * Custom response containing block transactions summary. Compatible with legacy fabric-socketio connector monitoring.
 *
 * @param cactiTransactionsEvents List of transactions summary
 */


data class CactiBlockTransactionsResponseV1 (

    /* List of transactions summary */
    @Json(name = "cactiTransactionsEvents")
    val cactiTransactionsEvents: kotlin.collections.List<CactiBlockTransactionEventV1>

)

