import React from "react";
import { StyleSheet, Text, TextInput, View } from "react-native";

// order matters here
import "react-native-polyfill-globals/auto";

import { StatusBar } from "expo-status-bar";
import { useGno } from "./src/hooks/use-gno";

// Polyfill async.Iterator. For some reason, the Babel presets and plugins are not doing the trick.
// Code from here: https://www.typescriptlang.org/docs/handbook/release-notes/typescript-2-3.html#caveats
(Symbol as any).asyncIterator =
  Symbol.asyncIterator || Symbol.for("Symbol.asyncIterator");

export default function App() {
  const gno = useGno();
  const [board, setBoard] = React.useState("");

  React.useEffect(() => {
    gno
      .setRemote("127.0.0.1:26557")
      .then((res) => {
        setBoard(res)
        console.warn("res ", JSON.stringify(res))
      })
      .catch((err) => setBoard("error " + JSON.stringify(err)));
  }, []);


  return (
    <View style={styles.container}>
      <Text>Flipppando</Text>
      <TextInput multiline={true} numberOfLines={40} value={board} />
      <StatusBar style="auto" />
    </View>
  );
}

const styles = StyleSheet.create({
  container: {
    flex: 1,
    backgroundColor: '#fff',
    alignItems: 'center',
    justifyContent: 'center',
  },
});
