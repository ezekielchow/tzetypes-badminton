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
import type { GameRecentStatistic } from './GameRecentStatistic';
import {
    GameRecentStatisticFromJSON,
    GameRecentStatisticFromJSONTyped,
    GameRecentStatisticToJSON,
} from './GameRecentStatistic';

/**
 * 
 * @export
 * @interface GetRecentStatistics200Response
 */
export interface GetRecentStatistics200Response {
    /**
     * 
     * @type {GameRecentStatistic}
     * @memberof GetRecentStatistics200Response
     */
    gameRecentStatistics: GameRecentStatistic;
}

/**
 * Check if a given object implements the GetRecentStatistics200Response interface.
 */
export function instanceOfGetRecentStatistics200Response(value: object): value is GetRecentStatistics200Response {
    if (!('gameRecentStatistics' in value) || value['gameRecentStatistics'] === undefined) return false;
    return true;
}

export function GetRecentStatistics200ResponseFromJSON(json: any): GetRecentStatistics200Response {
    return GetRecentStatistics200ResponseFromJSONTyped(json, false);
}

export function GetRecentStatistics200ResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): GetRecentStatistics200Response {
    if (json == null) {
        return json;
    }
    return {
        
        'gameRecentStatistics': GameRecentStatisticFromJSON(json['game_recent_statistics']),
    };
}

export function GetRecentStatistics200ResponseToJSON(value?: GetRecentStatistics200Response | null): any {
    if (value == null) {
        return value;
    }
    return {
        
        'game_recent_statistics': GameRecentStatisticToJSON(value['gameRecentStatistics']),
    };
}

