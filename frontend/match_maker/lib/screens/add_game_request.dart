import 'package:flutter/material.dart';
import 'package:http/http.dart' as http;
import 'dart:convert';
import 'package:match_maker/models/game_request.dart';
import 'package:match_maker/config.dart';
import 'package:syncfusion_flutter_datepicker/datepicker.dart';

class AddGameRequestScreen extends StatefulWidget {
  const AddGameRequestScreen({super.key});

  @override
  State<AddGameRequestScreen> createState() => _AddGameRequestScreenState();
}

class _AddGameRequestScreenState extends State<AddGameRequestScreen> {
  final _formKey = GlobalKey<FormState>();
  int _niveau = 0;
  String _location = 'Select a location';
  String _time = 'Select a time';
  String _gender = 'Any';
  int _amount = 0;
  double _price = 0;

  Future<void> _submitGameRequest() async {
    if (_formKey.currentState!.validate()) {
      _formKey.currentState!.save();

      final response = await http.post(
        Uri.parse('${Config.getBaseURL}/gamerequest'),
        headers: {'Content-Type': 'application/json'},
        body: jsonEncode(GameRequest(
          niveau: _niveau,
          location: _location,
          time: _time,
          gender: _gender,
          amount: _amount,
          price: _price,
        )),
      );

      if (response.statusCode == 201) {
        Navigator.pop(context, true); // Return to the previous screen
      } else {
        throw Exception('Failed to create game request');
      }
    }
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: const Text('Add Game Request'),
      ),
      body: Padding(
        padding: const EdgeInsets.all(16.0),
        child: Form(
          key: _formKey,
          child: ListView(
            children: <Widget>[
              TextFormField(
                decoration: const InputDecoration(labelText: 'Niveau'),
                keyboardType: TextInputType.number,
                onSaved: (value) => _niveau = int.parse(value!),
                validator: (value) {
                  if (value == null || value.isEmpty) {
                    return 'Please enter a niveau';
                  }
                  return null;
                },
              ),
              TextFormField(
                decoration: const InputDecoration(labelText: 'Location'),
                onSaved: (value) => _location = value!,
                validator: (value) {
                  if (value == null || value.isEmpty) {
                    return 'Please enter a location';
                  }
                  return null;
                },
              ),
              SfDateRangePicker(
                onSelectionChanged: (DateRangePickerSelectionChangedArgs args) {
                  _time = (args.value as DateTime).toIso8601String();
                },
              ),
              TextFormField(
                decoration: const InputDecoration(labelText: 'Time'),
                onSaved: (value) => _time = value!,
                validator: (value) {
                  if (value == null || value.isEmpty) {
                    return 'Please enter a time';
                  }
                  return null;
                },
              ),
              TextFormField(
                decoration: const InputDecoration(labelText: 'Gender'),
                onSaved: (value) => _gender = value!,
                validator: (value) {
                  if (value == null || value.isEmpty) {
                    return 'Please enter a gender';
                  }
                  return null;
                },
              ),
              TextFormField(
                decoration: const InputDecoration(labelText: 'Amount'),
                keyboardType: TextInputType.number,
                onSaved: (value) => _amount = int.parse(value!),
                validator: (value) {
                  if (value == null || value.isEmpty) {
                    return 'Please enter an amount';
                  }
                  return null;
                },
              ),
              TextFormField(
                decoration: const InputDecoration(labelText: 'Price'),
                keyboardType: TextInputType.numberWithOptions(decimal: true),
                onSaved: (value) => _price = double.parse(value!),
                validator: (value) {
                  if (value == null || value.isEmpty) {
                    return 'Please enter a price';
                  }
                  return null;
                },
              ),
              const SizedBox(height: 20),
              ElevatedButton(
                onPressed: _submitGameRequest,
                child: const Text('Submit'),
              ),
            ],
          ),
        ),
      ),
    );
  }
}
