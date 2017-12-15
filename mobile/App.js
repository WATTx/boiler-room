import React from 'react'
import { StyleSheet, View, Slider } from 'react-native'

import Client from './Client'

const client = new Client()

export default class App extends React.Component {

  onSlidingComplete = (val) => {
    client.setLevel("iot-room-6-north", "81c74c8f-5ae0-44a0-9164-6157c41564e0", val)
  }

  render() {
    return (
      <View style={styles.container}>
        <Slider
          style={{width: '50%'}}
          minimumValue={10}
          maximumValue={200}
          onSlidingComplete={this.onSlidingComplete} />
      </View>
    );
  }
}

const styles = StyleSheet.create({
  container: {
    flex: 1,
    backgroundColor: '#fff',
    alignItems: 'center',
    justifyContent: 'center',
  },
});
