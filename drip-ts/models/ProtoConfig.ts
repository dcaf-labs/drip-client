/* tslint:disable */
/* eslint-disable */
/**
 * Drip Backend
 * Drip backend service. All API\'s have a IP rate limit of 10 requests per second. 
 *
 * The version of the OpenAPI document: 1.0.0
 * Contact: mocha@dcaf.so
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
/**
 * 
 * @export
 * @interface ProtoConfig
 */
export interface ProtoConfig {
    /**
     * 
     * @type {string}
     * @memberof ProtoConfig
     */
    pubkey: string;
    /**
     * 
     * @type {string}
     * @memberof ProtoConfig
     */
    admin: string;
    /**
     * 
     * @type {string}
     * @memberof ProtoConfig
     */
    granularity: string;
    /**
     * 
     * @type {number}
     * @memberof ProtoConfig
     */
    tokenADripTriggerSpread: number;
    /**
     * 
     * @type {number}
     * @memberof ProtoConfig
     */
    tokenBWithdrawalSpread: number;
    /**
     * 
     * @type {number}
     * @memberof ProtoConfig
     */
    tokenBReferralSpread: number;
}

export function ProtoConfigFromJSON(json: any): ProtoConfig {
    return ProtoConfigFromJSONTyped(json, false);
}

export function ProtoConfigFromJSONTyped(json: any, ignoreDiscriminator: boolean): ProtoConfig {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'pubkey': json['pubkey'],
        'admin': json['admin'],
        'granularity': json['granularity'],
        'tokenADripTriggerSpread': json['tokenADripTriggerSpread'],
        'tokenBWithdrawalSpread': json['tokenBWithdrawalSpread'],
        'tokenBReferralSpread': json['tokenBReferralSpread'],
    };
}

export function ProtoConfigToJSON(value?: ProtoConfig | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'pubkey': value.pubkey,
        'admin': value.admin,
        'granularity': value.granularity,
        'tokenADripTriggerSpread': value.tokenADripTriggerSpread,
        'tokenBWithdrawalSpread': value.tokenBWithdrawalSpread,
        'tokenBReferralSpread': value.tokenBReferralSpread,
    };
}

