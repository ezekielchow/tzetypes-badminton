/* tslint:disable */
/* eslint-disable */
/**
 * tzetypes-badminton
 * tzetypes-badminton
 *
 * The version of the OpenAPI document: 1.0.0
 * 
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { mapValues } from '../runtime';
import type { InstagramMedia } from './InstagramMedia';
import {
    InstagramMediaFromJSON,
    InstagramMediaFromJSONTyped,
    InstagramMediaToJSON,
} from './InstagramMedia';

/**
 * 
 * @export
 * @interface GetInstagramFeed200Response
 */
export interface GetInstagramFeed200Response {
    /**
     * 
     * @type {Array<InstagramMedia>}
     * @memberof GetInstagramFeed200Response
     */
    feed: Array<InstagramMedia>;
}

/**
 * Check if a given object implements the GetInstagramFeed200Response interface.
 */
export function instanceOfGetInstagramFeed200Response(value: object): value is GetInstagramFeed200Response {
    if (!('feed' in value) || value['feed'] === undefined) return false;
    return true;
}

export function GetInstagramFeed200ResponseFromJSON(json: any): GetInstagramFeed200Response {
    return GetInstagramFeed200ResponseFromJSONTyped(json, false);
}

export function GetInstagramFeed200ResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): GetInstagramFeed200Response {
    if (json == null) {
        return json;
    }
    return {
        
        'feed': ((json['feed'] as Array<any>).map(InstagramMediaFromJSON)),
    };
}

export function GetInstagramFeed200ResponseToJSON(value?: GetInstagramFeed200Response | null): any {
    if (value == null) {
        return value;
    }
    return {
        
        'feed': ((value['feed'] as Array<any>).map(InstagramMediaToJSON)),
    };
}

