import 'dart:convert';

import 'package:flutter/material.dart';
import 'package:http/http.dart' as http;
import 'package:match_maker/models/game_request.dart';
import 'package:match_maker/config.dart';

Future<List<GameRequest>> fetchMatches() async {
  final response =
      await http.get(Uri.parse('http://localhost:8080/gamerequest'));

  if (response.statusCode == 200) {
    Iterable<dynamic> gameRequests = jsonDecode(response.body);
    return gameRequests.map((elem) => GameRequest.fromJson(elem)).toList();
  } else {
    throw Exception('Failed to load matches');
  }
}

void main() {
  runApp(const MyApp());
}

class MyApp extends StatelessWidget {
  const MyApp({super.key});

  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      title: 'MatchMaker',
      theme: ThemeData(
        colorScheme: ColorScheme.fromSeed(seedColor: Colors.deepPurple),
        useMaterial3: true,
      ),
      home: const MyHomePage(),
    );
  }
}

class MyHomePage extends StatefulWidget {
  const MyHomePage({super.key});

  @override
  State<MyHomePage> createState() => _MyHomePageState();
}

class _MyHomePageState extends State<MyHomePage> {
  late Future<List<GameRequest>> futureGameRequests;

  @override
  void initState() {
    super.initState();
    futureGameRequests = fetchMatches();
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        backgroundColor: Theme.of(context).colorScheme.inversePrimary,
        title: Text('MatchMaker'),
      ),
      body: Center(
        child: FutureBuilder<List<GameRequest>>(
          future: futureGameRequests,
          builder: (context, snapshot) {
            if (snapshot.connectionState == ConnectionState.waiting) {
              return const CircularProgressIndicator();
            } else if (snapshot.hasError) {
              return Text('Error: ${snapshot.error}');
            } else if (snapshot.hasData) {
              final gameRequests = snapshot.data!;
              return ListView.builder(
                itemCount: gameRequests.length,
                itemBuilder: (context, index) {
                  final gameRequest = gameRequests[index];
                  return ListTile(
                    title: Text(gameRequest.id.toString()),
                    subtitle: Text(gameRequest.time),
                  );
                },
              );
            } else {
              return const Text('No data');
            }
          },
        ),
      ),
    );
  }
}
