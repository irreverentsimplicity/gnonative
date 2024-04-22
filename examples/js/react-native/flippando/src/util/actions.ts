import AsyncStorage from '@react-native-async-storage/async-storage';
import {
  defaultFaucetTokenKey,
  defaultMnemonicKey,
} from '../types/types';
import { defaultTxFee, GnoJSONRPCProvider, GnoWallet, GnoWSProvider } from '@gnolang/gno-js-client';
import {
  BroadcastTxCommitResult,
  TM2Error,
  TransactionEndpoint
} from '@gnolang/tm2-js-client';
import { generateMnemonic } from './crypto';
import Long from 'long';
import Config from './config';
import { constructFaucetError } from './errors';

import {
  ErrorTransform
} from './errors';
import { UserFundedError } from '../types/errors';

// ENV values //
const wsURL: string = Config.GNO_WS_URL;
const rpcURL: string = Config.GNO_JSONRPC_URL;
const flippandoRealm: string = Config.GNO_FLIPPANDO_REALM;
const faucetURL: string = "http://127.0.0.1:5050";
const defaultGasWanted: Long = new Long(1000_000_0);
const customTXFee = '2000000ugnot'

const cleanUpRealmReturn = (ret: string) => {
  return ret.slice(2, -9).replace(/\\"/g, '"');
};
const decodeRealmResponse = (resp: string) => {
  return cleanUpRealmReturn(atob(resp));
};
const parsedJSONOrRaw = (data: string, nob64 = false) => {
  const decoded = nob64 ? cleanUpRealmReturn(data) : decodeRealmResponse(data);
  try {
    return JSON.parse(decoded);
  } finally {
    return decoded;
  }
};

class Actions {
  private static instance: Actions;
  private static initPromise: Actions | PromiseLike<Actions>;
  private wallet: GnoWallet | null = null;
  private provider: GnoWSProvider | null = null;
  private providerJSON: GnoJSONRPCProvider | null = null;
  private faucetToken: string | null = null;
  
  private constructor() {}

  public static async getInstance(): Promise<Actions> {
    if (!Actions.instance) {
      Actions.instance = new Actions();
      Actions.initPromise = new Promise(async (resolve) => {
        await Actions.instance.initialize();
        resolve(Actions.instance);
      });
      return Actions.initPromise;
    } else {
      return Actions.initPromise;
    }
  }

  private async initialize() {
    let mnemonic: string | null = await AsyncStorage.getItem(defaultMnemonicKey);
    if (!mnemonic || mnemonic === '') {
      mnemonic = generateMnemonic();
      await AsyncStorage.setItem(defaultMnemonicKey, mnemonic);
    }
    try {
      this.wallet = await GnoWallet.fromMnemonic(mnemonic);
      this.providerJSON = new GnoJSONRPCProvider(rpcURL);
      this.wallet.connect(this.providerJSON);
    } catch (e) {
      console.error('Could not create wallet from mnemonic');
    }

    let faucetToken: string | null = await AsyncStorage.getItem(defaultFaucetTokenKey);
    if (faucetToken && faucetToken !== '') {
      this.faucetToken = faucetToken;
      try {
        await this.fundAccount(this.faucetToken);
      } catch (e) {
        if (e instanceof UserFundedError) {
          console.log('User already funded.');
        } else {
          console.log('Could not fund user.');
        }
      }
    }
  }

  public async setFaucetToken(token: string) {
    await this.fundAccount(token);
    this.faucetToken = token;
    await AsyncStorage.setItem(defaultFaucetTokenKey, token);
  }

  public async getFaucetToken(): Promise<string | null> {
    return this.faucetToken || await AsyncStorage.getItem(defaultFaucetTokenKey);
  }

  // Implement additional methods following the pattern above

  public destroy() {
    if (!this.provider) {
      return;
    }
    // Close any necessary connections or cleanup if needed
  }

  private async fundAccount(token: string): Promise<void> {
    const requestOptions = {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
        'faucet-token': 'flippando'
      },
      body: JSON.stringify({
        to: await this.wallet?.getAddress()
      })
    };

    await fetch(faucetURL, requestOptions).then(fResponse => {
      console.log("faucetURL", faucetURL)
      console.log("fundResponse", JSON.stringify(fResponse, null, 2))
      if (!fResponse.ok) {
        console.log("fund error, ", fResponse.text());
      }
    });
  }
}

export default Actions;

