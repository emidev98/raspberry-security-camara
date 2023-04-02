import { Moment } from 'moment';
import {useCookies} from 'react-cookie';

const COOIKE_NAME = 'authToken';

export const useCookie = () => {
    const [cookies, setCookie, removeCookie] = useCookies([COOIKE_NAME]);

    function setToken(token: string, expiration?: Moment) {
        let expires = expiration ? expiration.toDate() : undefined;

        setCookie(COOIKE_NAME, token, { 
            expires: expires, 
            secure: true, 
            sameSite: 'strict'
        });
    }

    const getToken = () => cookies?.authToken || '';

    const rmToken = () => removeCookie(COOIKE_NAME);

    return { setToken, getToken, rmToken }
}