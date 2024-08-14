import 'dart:io' show Platform;

class Config {
  static const String _localHostURI = 'http://localhost:8080';
  static const String _emulatorURI = 'http://10.0.2.2:8080';

  static String getBaseURL =
      Platform.isWindows ? _localHostURI :
      Platform.isAndroid ? _emulatorURI : 
      Platform.isIOS ? _emulatorURI :
      throw Exception('Unknown platform');
}