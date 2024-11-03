import 'dart:convert';
import 'dart:io';
import 'package:http/http.dart' as http;
import 'package:flutter/material.dart';
import 'package:table_calendar/table_calendar.dart';
import 'package:match_maker/config.dart';
import 'package:match_maker/models/game_request.dart';
import 'add_game_request.dart';

Future<List<GameRequest>> fetchMatches(DateTime date) async {
  try {
    final response = await http.get(Uri.parse(
        '${Config.getBaseURL}/gamerequests?date=${date.toIso8601String()}'));

    if (response.statusCode == 200) {
      Iterable<dynamic> gameRequests = jsonDecode(response.body);
      return gameRequests.map((elem) => GameRequest.fromJson(elem)).toList();
    } else {
      throw Exception('Failed to load matches');
    }
  } catch (e) {
    // Log the exception or handle it gracefully
    if (e is SocketException) {
      // Server is unreachable
      print('Server is unreachable: $e');
    } else {
      print('An error occurred: $e');
    }
    // Return an empty list or handle the error appropriately
    return [];
  }
}

class MyHomePage extends StatefulWidget {
  const MyHomePage({super.key});

  @override
  State<MyHomePage> createState() => _MyHomePageState();
}

class _MyHomePageState extends State<MyHomePage> {
  late Future<List<GameRequest>> futureGameRequests;
  DateTime _selectedDay = DateTime.now();
  CalendarFormat _calendarFormat = CalendarFormat.week;

  @override
  void initState() {
    super.initState();
    futureGameRequests = fetchMatches(_selectedDay);
  }

  Future<void> _navigateToAddGameRequest() async {
    final result = await Navigator.push(
      context,
      MaterialPageRoute(builder: (context) => const AddGameRequestScreen()),
    );

    if (result == true) {
      // Refresh the list of game requests
      setState(() {
        futureGameRequests = fetchMatches(_selectedDay);
      });
    }
  }

  void _onDaySelected(DateTime selectedDay, DateTime focusedDay) {
    setState(() {
      _selectedDay = selectedDay;
      futureGameRequests = fetchMatches(_selectedDay);
    });
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        backgroundColor: Theme.of(context).colorScheme.inversePrimary,
        title: const Text('MatchMaker'),
      ),
      body: Column(
        children: [
          TableCalendar(
            firstDay: DateTime.utc(2000, 1, 1),
            lastDay: DateTime.utc(2100, 1, 1),
            focusedDay: DateTime.now(),
            onDaySelected: _onDaySelected,
            calendarFormat: _calendarFormat,
            selectedDayPredicate: (day) {
              return isSameDay(_selectedDay, day);
            },
            onFormatChanged: (format) {
              print(format);
              setState(() {
                _calendarFormat = format;
              });
            },
          ),
          Expanded(
            child: FutureBuilder<List<GameRequest>>(
              future: futureGameRequests,
              builder: (context, snapshot) {
                if (snapshot.connectionState == ConnectionState.waiting) {
                  return const Center(
                    child: CircularProgressIndicator(),
                  );
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
                        subtitle: Text(gameRequest.time.toString()),
                      );
                    },
                  );
                } else {
                  return const Text('No data');
                }
              },
            ),
          ),
        ],
      ),
      floatingActionButton: FloatingActionButton(
        onPressed: _navigateToAddGameRequest,
        child: const Icon(Icons.add),
      ),
    );
  }
}
