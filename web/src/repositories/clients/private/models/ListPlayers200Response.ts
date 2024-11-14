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
import type { Player } from './Player';
import {
    PlayerFromJSON,
    PlayerFromJSONTyped,
    PlayerToJSON,
} from './Player';
import type { ListPlayers200ResponsePagination } from './ListPlayers200ResponsePagination';
import {
    ListPlayers200ResponsePaginationFromJSON,
    ListPlayers200ResponsePaginationFromJSONTyped,
    ListPlayers200ResponsePaginationToJSON,
} from './ListPlayers200ResponsePagination';

/**
 * 
 * @export
 * @interface ListPlayers200Response
 */
export interface ListPlayers200Response {
    /**
     * 
     * @type {Array<Player>}
     * @memberof ListPlayers200Response
     */
    players?: Array<Player>;
    /**
     * 
     * @type {ListPlayers200ResponsePagination}
     * @memberof ListPlayers200Response
     */
    pagination?: ListPlayers200ResponsePagination;
}

/**
 * Check if a given object implements the ListPlayers200Response interface.
 */
export function instanceOfListPlayers200Response(value: object): value is ListPlayers200Response {
    return true;
}

export function ListPlayers200ResponseFromJSON(json: any): ListPlayers200Response {
    return ListPlayers200ResponseFromJSONTyped(json, false);
}

export function ListPlayers200ResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): ListPlayers200Response {
    if (json == null) {
        return json;
    }
    return {
        
        'players': json['players'] == null ? undefined : ((json['players'] as Array<any>).map(PlayerFromJSON)),
        'pagination': json['pagination'] == null ? undefined : ListPlayers200ResponsePaginationFromJSON(json['pagination']),
    };
}

export function ListPlayers200ResponseToJSON(value?: ListPlayers200Response | null): any {
    if (value == null) {
        return value;
    }
    return {
        
        'players': value['players'] == null ? undefined : ((value['players'] as Array<any>).map(PlayerToJSON)),
        'pagination': ListPlayers200ResponsePaginationToJSON(value['pagination']),
    };
}

